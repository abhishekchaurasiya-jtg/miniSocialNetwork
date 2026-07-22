package routes

import (
	gin "github.com/gin-gonic/gin"

	controllers "app/src/controllers"
)


func RegisterPublicAuthRoutes(router gin.IRouter, authController *controllers.AuthController) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", authController.SignUp)
		auth.POST("/login", authController.Login)
		auth.POST("/refresh", authController.RefreshToken)
	}
}


func RegisterPrivateAuthRoutes(router gin.IRouter, authController *controllers.AuthController) {
	auth := router.Group("/auth")
	{
		auth.POST("/update_password", authController.UpdatePassword)
		auth.POST("/logout", authController.LogOut)
	}
}