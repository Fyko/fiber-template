package errors

import "github.com/gofiber/fiber/v2"

func CreateReSTError(message string) fiber.Map {
	return fiber.Map{
		"success": false,
		"message": message,
	}
}
