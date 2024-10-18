package routes

import (
	"github.com/dzakimaulana/SiJaki-Backend/internal/handlers"
	"github.com/dzakimaulana/SiJaki-Backend/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(h *handlers.UserHandler, r fiber.Router) {
	r.Post("/login", h.Login)
	r.Post("/register", h.Register)
	r.Post("/logout", middlewares.OnlyAdmin, h.Logout)
}
