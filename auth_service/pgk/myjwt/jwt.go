package myjwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(email string) (string, error) {
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(60 * time.Hour).Unix(),
		"sub": email,
	})

	signedToken, err := rawToken.SignedString("1234")
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
