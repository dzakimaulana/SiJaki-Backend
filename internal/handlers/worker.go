package handlers

import (
	"strconv"

	"github.com/dzakimaulana/SiJaki-Backend/internal/models"
	"github.com/dzakimaulana/SiJaki-Backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

type WorkerHandler struct {
	workerSvc *services.WorkerSvc
}

func NewWorkerHandler(ws *services.WorkerSvc) *WorkerHandler {
	return &WorkerHandler{
		workerSvc: ws,
	}
}

func (wh *WorkerHandler) AddWorker(c *fiber.Ctx) error {
	var worker models.Worker

	if err := c.BodyParser(&worker); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Failed to parse request",
		})
	}

	if err := wh.workerSvc.AddWorker(&worker); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Failed to add worker data",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Add Worker Success",
	})
}

func (wh *WorkerHandler) EditWorker(c *fiber.Ctx) error {
	var worker models.Worker

	if err := c.BodyParser(&worker); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Failed to parse request",
		})
	}

	existingWorker, err := wh.workerSvc.GetWorkerByID(worker.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Failed to find worker",
		})
	}
	if existingWorker == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  "Worker not found",
		})
	}

	if err := wh.workerSvc.EditWorker(&worker); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Failed to edit worker",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Edit Worker Success",
	})
}

func (wh *WorkerHandler) DeleteWorker(c *fiber.Ctx) error {
	workerIdStr := c.Params("id")
	workerId, err := strconv.ParseUint(workerIdStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid worker ID format",
		})
	}

	existingWorker, err := wh.workerSvc.GetWorkerByID(uint(workerId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Failed to search for worker",
		})
	}

	if existingWorker == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  "Worker not found",
		})
	}

	if err := wh.workerSvc.DeleteWorker(uint(workerId)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Failed to delete worker",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "success",
		"message": "Worker deleted successfully",
	})
}

func (wh *WorkerHandler) GetAllWorkers(c *fiber.Ctx) error {
	allWorkers, err := wh.workerSvc.GetAllWorkers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  "Failed to get all workers data",
		})
	}

	message := "Get Worker Success"
	if len(allWorkers) == 0 {
		message = "No workers added"
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": message,
		"data":    allWorkers,
	})
}
