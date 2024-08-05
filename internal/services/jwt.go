package services

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString(SecretKey)
}
