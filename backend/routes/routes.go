package routes

import (
	"AST-Generator/config"
	"AST-Generator/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/parse", controllers.ParseCode)

	router.GET("/auth/google/login", func(c *gin.Context) {
		url := config.GoogleOauthConfig.AuthCodeURL("str", oauth2.AccessTypeOffline)
		c.Redirect(http.StatusTemporaryRedirect, url)
	})
	router.GET("/auth/google/callback", controllers.GoogleCallbackHandler)
}
