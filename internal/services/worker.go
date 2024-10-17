package services

import (
	"fmt"

	"github.com/dzakimaulana/SiJaki-Backend/internal/models"
	"gorm.io/gorm"
)

type WorkerSvc struct {
	DB *gorm.DB
}

func NewWorkerSvc(db *gorm.DB) *WorkerSvc {
	return &WorkerSvc{
		DB: db,
	}
}

func (ws *WorkerSvc) AddWorker(worker *models.Worker) error {
	if err := ws.DB.Create(worker).Error; err != nil {
		return err
	}
	return nil
}

func (ws *WorkerSvc) EditWorker(worker *models.Worker) error {
	var existingWorker models.Worker
	if err := ws.DB.First(&existingWorker, worker.ID).Error; err != nil {
		return err
	}

	if err := ws.DB.Model(&existingWorker).Updates(worker).Error; err != nil {
		return err
	}

	return nil
}

func (ws *WorkerSvc) DeleteWorkerByID(id uint) error {
	var worker models.Worker
	if err := ws.DB.First(&worker, id).Error; err != nil {
		return err
	}

	if err := ws.DB.Delete(&worker).Error; err != nil {
		return err
	}

	return nil
}

func (ws *WorkerSvc) GetWorkerByID(id uint) (*models.Worker, error) {
	var worker models.Worker

	if err := ws.DB.First(&worker, id).Error; err != nil {
		return nil, fmt.Errorf("worker not found: %w", err)
	}

	return &worker, nil
}
