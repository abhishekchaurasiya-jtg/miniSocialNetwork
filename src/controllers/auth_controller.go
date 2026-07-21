package controllers

import (
	http "net/http"
	
	gin "github.com/gin-gonic/gin"

	dto "app/src/dto"
	services "app/src/services"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func SetCookieToContext(context *gin.Context, tokens dto.TokensCollection) {
	var age int = 60 * 24 * 30 * 2 // Seconds * Hours * Days * Months  // 2months(60days)
	context.SetCookie("refresh_token", tokens.RefreshToken, age, "/auth/refresh", "localhost", false, false)
}


func (authCnt *AuthController) SignUp(c *gin.Context) {
	var request dto.CreateUserRequest

	// Bind and validate request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Bussiness Logic
	tokens, err := authCnt.authService.Register(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	SetCookieToContext(c, *tokens)
	// Success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"token":   tokens.AccessToken,
	})
}

func (authCnt *AuthController) Login(c *gin.Context) {
	var request dto.LoginUserRequest

	// Bind and validate request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tokens, err := authCnt.authService.Login(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	SetCookieToContext(c, *tokens)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"token":   tokens.AccessToken,
	})
}