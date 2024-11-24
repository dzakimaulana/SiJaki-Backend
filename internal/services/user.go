package services

import (
	"errors"

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

	if err := us.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
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
	if err := us.DB.Where("id = ?", user.ID).First(&existingUser).Error; err != nil {
		return err
	}

	if err := us.DB.Model(&existingUser).Updates(user).Error; err != nil {
		return err
	}

	return nil
}
