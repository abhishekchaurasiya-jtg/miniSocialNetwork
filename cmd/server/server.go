package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.Run()
}
