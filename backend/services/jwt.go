package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID, email, name, picture, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{
		"sub":     userID,
		"email":   email,
		"name":    name,
		"picture": picture,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
