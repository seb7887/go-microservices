package auth

import (
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/seb7887/go-microservices/config"
)

func ValidateToken(jwtToken string) (string, bool) {
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JWTSecret), nil
	})
	if err != nil {
		return "", false
	}
	if token.Valid {
		userId := tokenData["user_id"].(string)
		return userId, true
	} else {
		return "Invalid Token", false
	}
}