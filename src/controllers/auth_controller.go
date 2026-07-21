package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	services "app/src/services"
	dto "app/src/dto"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}


func (a *AuthController) SignUp(c *gin.Context) {
	var req dto.CreateUserRequest

	// Bind and validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Bussiness Logic
	token, err := a.authService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"token":   token,
	})
}