package server

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"

	config "app/config"
	db_pool "app/db"
	controllers "app/src/controllers"
	repositories "app/src/repositories"
	routes "app/src/router"
	services "app/src/services"
)

func healthEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "alive",
	})
}

func RunServer() {
	cnf := config.LoadConfig()
	_, gormDB := db_pool.InitDB(cnf)
	jwtService := services.GetJwtService([]byte(cnf.JWTSecretKey), cnf.JWTIssuer)

	router := gin.Default()
	// Public routes
	public := router.Group("")
	{
		// Health API
		public.GET("/health", healthEndpoint)

		// Authentication API
		userRepo := repositories.NewUserRepository(gormDB)
		authService := services.NewAuthService(userRepo, jwtService)
		authController := controllers.NewAuthController(authService)
		routes.RegisterAuthRoutes(public, authController)
	}

	router.Run()
}
