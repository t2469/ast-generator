package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"AST-Generator/config"
	"AST-Generator/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GoogleCallbackHandler(c *gin.Context) {
	stateQuery := c.Query("state")
	if stateQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "State not found in query"})
		return
	}

	session := sessions.Default(c)
	storedState := session.Get("oauthState")
	if storedState == nil || stateQuery != storedState {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state parameter"})
		return
	}
	session.Delete("oauthState")
	session.Save()

	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
		return
	}
	token, err := config.GoogleOauthConfig.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Token exchange failed: %s", err.Error())})
		return
	}
	client := config.GoogleOauthConfig.Client(c, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get user info: %s", err.Error())})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read user info"})
		return
	}

	var userInfo map[string]interface{}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
		return
	}

	userID, _ := userInfo["sub"].(string)
	email, _ := userInfo["email"].(string)

	jwtToken, err := services.GenerateJWT(userID, email, config.JwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
		"user":  userInfo,
	})
}
