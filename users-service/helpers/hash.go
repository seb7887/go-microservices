package helpers

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashed)
}

func VerifyPassword(password string, hash string) bool {
	passErr := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return false
	}

	return true
}