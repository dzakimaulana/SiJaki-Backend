package database

import (
	"fmt"
	"log"

	"github.com/dzakimaulana/SiJaki-Backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(user string, dbname string, password string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("😞 Failed to connect to database: %w", err)
	}

	// Auto migrate your models
	err = db.AutoMigrate(&models.User{}, &models.Worker{})
	if err != nil {
		log.Fatalf("😞 Failed to migrate database: %v", err)
		return nil, err
	}

	return db, nil
}
