package routes

import (
	gin "github.com/gin-gonic/gin"

	controllers "app/src/controllers"
)



func RegisterPrivateUserRoutes(router gin.IRouter, userController *controllers.UserController) {
	userRouter := router.Group("/user")
	{
		userRouter.PATCH("/update", userController.UpdateUser)
		userRouter.DELETE("/delete", userController.DeleteUser)
		userRouter.GET("/me", userController.GetUser)
		userRouter.GET("/all", userController.GetActiveUsers)  // query params will be passed
	}
}