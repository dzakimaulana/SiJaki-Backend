package utils

import (
	"fmt"
	"time"

	"github.com/dzakimaulana/SiJaki-Backend/internal/config"
	"github.com/dzakimaulana/SiJaki-Backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

var cfgJWT = config.LoadJWTConfig()

func GenerateJWT(user *models.User) (string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)

	claims := jwt.MapClaims{
		"exp":        expirationTime.Unix(),
		"authorized": true,
		"user":       user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfgJWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfgJWT.Secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token or claims")
}
