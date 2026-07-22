package server

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"

	config "app/config"
	db_pool "app/db"
	controllers "app/src/controllers"
	middlewares "app/src/middlewares"
	repositories "app/src/repositories"
	routes "app/src/router"
	services "app/src/services"
)

func healthEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
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
		routes.RegisterPublicAuthRoutes(public, authController)
	}

	protected := router.Group("")
	protected.Use(middlewares.AuthTokenMiddleware(jwtService))
	{
		// Authentication Private Routes
		userRepo := repositories.NewUserRepository(gormDB)
		authService := services.NewAuthService(userRepo, jwtService)
		authController := controllers.NewAuthController(authService)
		routes.RegisterPrivateAuthRoutes(protected, authController)

		// Remaining CRUD API's
	}

	router.Run()
}
