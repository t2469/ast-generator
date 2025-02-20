package routes

import (
	"AST-Generator/config"
	"AST-Generator/controllers"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/parse", controllers.ParseCode)

	router.GET("/auth/google/login", func(c *gin.Context) {
		state := config.GenerateState(16)

		session := sessions.Default(c)
		session.Set("oauthState", state)
		session.Save()

		url := config.GoogleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
		c.Redirect(http.StatusTemporaryRedirect, url)
	})
	router.GET("/auth/logout", controllers.LogoutHandler)

	router.GET("/auth/google/callback", controllers.GoogleCallbackHandler)
	router.GET("/auth/current_user", controllers.JWTAuthMiddleware(), controllers.GetCurrentUser)
}
