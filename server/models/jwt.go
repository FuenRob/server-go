package model

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func GenerateUserToken(name string) string {
	// Set custom claims
	claims := &jwtCustomClaims{
		name,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	userToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err.Error()
	}

	return userToken
}
