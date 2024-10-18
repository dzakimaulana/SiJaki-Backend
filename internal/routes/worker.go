package routes

import (
	"github.com/dzakimaulana/SiJaki-Backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func WorkerRoute(h *handlers.WorkerHandler, r fiber.Router) {
	r.Get("/", h.GetAllWorkers)
	r.Post("/add", h.AddWorker)
	r.Put("/edit", h.EditWorker)
	r.Delete("/delete/:id", h.DeleteWorker)
}
