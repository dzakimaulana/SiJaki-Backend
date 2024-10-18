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

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Failed to parse request",
		})
	}

	existingUser, err := uh.userSvc.GetUserByUsername(user.Username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Check username & password",
		})
	}

	correctUser := utils.CheckPassword(user.Password, existingUser.Password)
	if !correctUser {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Check username & password",
		})
	}

	token, err := utils.GenerateJWT(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error when generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login Succes",
		"data": fiber.Map{
			"token": token,
		},
	})
}

func (uh *UserHandler) Register(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Failed to parse request",
		})
	}

	existingUser, err := uh.userSvc.GetUserByUsername(user.Username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Username Already Used",
		})
	}
	if existingUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Username Already Used",
		})
	}

	if err := uh.userSvc.AddUser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Failed add user to database",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Register Success",
	})
}

func (uh *UserHandler) Logout(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Logout Success",
	})
}
