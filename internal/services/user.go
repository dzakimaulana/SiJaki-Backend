package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dzakimaulana/SiJaki-Backend/internal/models"
	"gorm.io/gorm"
)

type UserSvc struct {
	DB *gorm.DB
}

func NewUserSvc(db *gorm.DB) *UserSvc {
	return &UserSvc{
		DB: db,
	}
}

func (us *UserSvc) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := us.DB.First(&user, username).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *UserSvc) AddUser(user *models.User) error {
	if err := us.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (us *UserSvc) EditUser(user *models.User) error {
	var existingUser models.User
	if err := us.DB.First(&existingUser, user.ID).Error; err != nil {
		return err
	}

	if err := us.DB.Model(&existingUser).Updates(user).Error; err != nil {
		return err
	}

	return nil
}

func (ws *UserSvc) GenerateJWT(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("rahasia")) // change with .env
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
