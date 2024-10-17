package handlers

import (
	"github.com/dzakimaulana/SiJaki-Backend/internal/models"
	"github.com/dzakimaulana/SiJaki-Backend/internal/services"
	"github.com/dzakimaulana/SiJaki-Backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userSvc *services.UserSvc
}

func NewUserHandler(us *services.UserSvc) *UserHandler {
	return &UserHandler{
		userSvc: us,
	}
}

func (uh *UserHandler) Login(c *fiber.Ctx) error {
	var user models.User

	// Parse the incoming JSON payload
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request",
		})
	}

	existingUser, err := uh.userSvc.GetUserByUsername(user.Username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Check username & password",
		})
	}

	correctUser := utils.CheckPassword(user.Password, existingUser.Password)
	if !correctUser {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Check username & password",
		})
	}

	token, err := uh.userSvc.GenerateJWT(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error when generate token",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Login Succes",
		"token":   token,
	})
}
