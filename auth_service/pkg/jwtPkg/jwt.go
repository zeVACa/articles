package jwtPkg

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("a-string-secret-at-least-256-bits-long")

type Claims struct {
	Email  string `json:"user_email"`
	UserID int64  `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string) (string, error) {
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_email": email,
		"iat":        time.Now().Unix(),
		"exp":        time.Now().Add(60 * time.Hour).Unix(),
	})

	signedToken, err := rawToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken - Единственный код, который взял из нейронки
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
