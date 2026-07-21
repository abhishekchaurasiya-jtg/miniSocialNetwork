package routes

import (
	gin "github.com/gin-gonic/gin"

	controllers "app/src/controllers"
)

func RegisterAuthRoutes(router *gin.Engine, authController *controllers.AuthController) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", authController.SignUp)

		// Future APIs
		// auth.POST("/login", authController.Login)
		// auth.POST("/refresh", authController.RefreshToken)
		// auth.POST("/logout", authController.Logout)
	}
}