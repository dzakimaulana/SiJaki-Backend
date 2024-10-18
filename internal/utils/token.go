package utils

import (
	"fmt"
	"time"

	"github.com/dzakimaulana/SiJaki-Backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user *models.User) (string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)

	claims := jwt.MapClaims{
		"exp":        expirationTime.Unix(),
		"authorized": true,
		"user":       user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte("rahasia")

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*jwt.MapClaims, error) {
	secretKey := []byte("rahasia")

	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token or claims")
}
