package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type TokensCollection struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JWTService interface {
	GenerateNewTokens(userID int, email string) (*TokensCollection, error)
	GenerateAccessTokenFromRefresh(refreshTokenString string) (string, error)
	ValidateToken(tokenString string) (*UserClaims, error)
}

type UserClaims struct {
	UserId int    `json:"user_id"` 
	Email  string `json:"email"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey []byte
	issuer    string
}

func NewJWTService(secret string, issuer string) JWTService {
	return &jwtService{
		secretKey: []byte(secret),
		issuer:    issuer,
	}
}

func (j *jwtService) GenerateNewTokens(userID int, email string) (*TokensCollection, error) {
	accessTokenClaims := &UserClaims{
		UserId:    userID,
		Email:     email,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 4)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.issuer,
		},
	}

	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessToken, err := accessTokenObj.SignedString(j.secretKey)
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := &UserClaims{
		UserId:    userID,
		Email:     email,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.issuer,
		},
	}
	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshToken, err := refreshTokenObj.SignedString(j.secretKey)
	if err != nil {
		return nil, err
	}

	return &TokensCollection{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (j *jwtService) GenerateAccessTokenFromRefresh(refreshTokenString string) (string, error) {
	claims, err := j.ValidateToken(refreshTokenString)
	if err != nil {
		return "", err
	}

	if claims.TokenType != "refresh" {
		return "", errors.New("provided token is not a valid refresh token type")
	}


	newAccessClaims := &UserClaims{
		UserId:    claims.UserId,
		Email:     claims.Email,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.issuer,
		},
	}

	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, newAccessClaims)
	return accessTokenObj.SignedString(j.secretKey)
}

func (j *jwtService) ValidateToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected token cryptographic signing method")
		}
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token payload parsing failed or context expired")
}
