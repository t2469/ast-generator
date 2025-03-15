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

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.GET("/auth/google/login", func(c *gin.Context) {
		state := config.GenerateState(16)

		session := sessions.Default(c)
		session.Set("oauthState", state)
		session.Save()

		url := config.GoogleOauthConfig.AuthCodeURL(
			state,
			oauth2.AccessTypeOffline,
			oauth2.SetAuthURLParam("prompt", "select_account"),
			oauth2.SetAuthURLParam("include_granted_scopes", "true"),
		)
		c.Redirect(http.StatusTemporaryRedirect, url)
	})
	router.GET("/auth/google/callback", controllers.GoogleCallbackHandler)
	router.GET("/auth/logout", controllers.LogoutHandler)
	router.GET("/auth/current_user", controllers.JWTAuthMiddleware(), controllers.GetCurrentUser)

	router.GET("/source_codes", controllers.GetAllSourceCodesHandler)
	router.GET("/source_codes/user", controllers.JWTAuthMiddleware(), controllers.GetUserSourceCodesHandler)
	router.POST("/source_codes/save", controllers.JWTAuthMiddleware(), controllers.SaveSourceCodeHandler)
	router.DELETE("/source_codes/:id", controllers.JWTAuthMiddleware(), controllers.DeleteSourceCodeHandler)
}
