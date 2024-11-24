package middlewares

import (
	"strings"

	"github.com/dzakimaulana/SiJaki-Backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func OnlyAdmin(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if len(token) > 7 && strings.ToLower(token[:7]) == "bearer " {
		token = token[7:]
	}

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "No token provided",
		})
	}

	claims, err := utils.VerifyJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to verify token",
		})
	}

	user, ok := claims["user"].(string)
	if !ok || user != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Only admin access allowed",
		})
	}

	c.Locals("claims", claims)

	return c.Next()
}
