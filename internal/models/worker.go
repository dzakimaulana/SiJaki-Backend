package models

type Worker struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Age         uint   `json:"age" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}
