package middlewares

import (
	"log"
	http "net/http"
	strings "strings"

	gin "github.com/gin-gonic/gin"

	services "app/src/services"
)


func AuthTokenMiddleware(jwtService services.JWTService) gin.HandlerFunc {
	return func(context *gin.Context) {

		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusUnauthorized,  gin.H{"error": "Authorization header is missing"})
			context.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			context.Abort()
			return
		}

		tokenString := parts[1]

		userClaims, err := jwtService.ValidateToken(tokenString) 
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or empty token",
			})
			context.Abort()
			return 
		}

		log.Println("AuthMiddleware: UserToken Verified.")
		log.Println("AuthMiddleware: Recieved Values", userClaims.UserId, userClaims.Email)

		context.Set("UserID", userClaims.UserId)
		context.Set("Email", userClaims.Email)

		context.Next()
	}
}