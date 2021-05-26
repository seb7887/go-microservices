package helpers

import (
	"log"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/seb7887/go-microservices/config"
)

func GenerateJWT(userId string, email string) string {
	tokenContent := jwt.MapClaims{
		"user_id": userId,
		"email": email,
		"expiry": time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte(config.GetConfig().JWTSecret))
	if err != nil {
		log.Fatal(err)
	}

	return token
}