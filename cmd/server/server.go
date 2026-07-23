package server

import (
	http "net/http"

	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"

	_ "app/src/dto/custom_validators"
	config "app/config"
	controllers "app/src/controllers"
	db_pool "app/db"
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
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Enables cookies 
	}))
	
	vrouter := router.Group("/api/v1")
	// Public routes
	public := vrouter.Group("")
	{
		// Health API
		public.GET("/health", healthEndpoint)

		// Authentication API
		authRepo := repositories.NewAuthRepository(gormDB)
		authService := services.NewAuthService(authRepo, jwtService)
		authController := controllers.NewAuthController(authService)
		routes.RegisterPublicAuthRoutes(public, authController)
	}

	protected := vrouter.Group("")
	protected.Use(middlewares.AuthTokenMiddleware(jwtService))
	{
		// Authentication Private Routes
		authRepo := repositories.NewAuthRepository(gormDB)
		authService := services.NewAuthService(authRepo, jwtService)
		authController := controllers.NewAuthController(authService)
		routes.RegisterPrivateAuthRoutes(protected, authController)

		userRepo := repositories.NewUserRepository(gormDB)
		userService := services.NewUserService(userRepo)
		userController := controllers.NewUserController(userService)
		routes.RegisterPrivateUserRoutes(protected, userController)
	}

	router.Run()
}
