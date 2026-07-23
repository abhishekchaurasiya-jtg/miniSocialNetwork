package server

import (
	http "net/http"
	time "time"

	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	// CORS Configuration
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "PATCH", "POST", "DELETE"},
		AllowHeaders: []string{"Origin", "Authorization"},
		ExposeHeaders:[]string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.Run()
}
