package controllers

import (
	http "net/http"
	"time"

	gin "github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"

	dto "app/src/dto"
	services "app/src/services"
)


var validate = validator.New()

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
	responsePayload, tokens, err := authCnt.authService.Register(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := validate.Struct(responsePayload); err != nil {  // Explore
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Schema Error: Outbound data daes not match the required schema format",
		})
		return
	}
	
	SetCookieToContext(c, *tokens)
	// Success response
	c.JSON(http.StatusCreated, responsePayload)
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
	loginPayload, tokens, err := authCnt.authService.Login(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := validate.Struct(loginPayload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Schema Error: Outbound data daes not match the required schema format",
		})
		return
	}
	SetCookieToContext(c, *tokens)
	c.JSON(http.StatusOK, loginPayload)
}


func (authCnt *AuthController) RefreshToken(context *gin.Context) {
	cookieValue, err := context.Cookie("refresh_token")

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authentication cookie missing or expired",
		})
		return
	}
	accessToken, err := authCnt.authService.RefreshAccessToken(cookieValue)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// returning the response
	context.JSON(
		http.StatusCreated, dto.TokenResponse{
			Token: *accessToken,
			ExpiryTime: time.Now().Add(time.Minute*15),
		},
	)
}

func (authCnt *AuthController) UpdatePassword(context *gin.Context) {
	var request *dto.UpdatePasswordRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	
	email, exists := context.Get("Email")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Email not found in context"})
		return
	}

	emailStr, ok := email.(string)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Email context value is not a valid string"})
		return
	}
	err := authCnt.authService.UpdatePassword(*request, emailStr)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	context.Status(http.StatusNoContent)
}

func (authCnt *AuthController) LogOut(context *gin.Context) {

	context.SetCookie("refresh_token", "", -1, "/auth/refresh", "localhost", false, false)
	email, exists := context.Get("Email")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Email not found in context"})
		return
	}

	emailStr, ok := email.(string)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Email context value is not a valid string"})
		return
	}

	err := authCnt.authService.LogOut(emailStr)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully. Please delete token from client storage."},
	)
}


