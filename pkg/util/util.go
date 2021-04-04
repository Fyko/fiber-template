package util

import "github.com/gofiber/fiber/v2"

func CreateReSTResponse(success bool, message interface{}) fiber.Map {
	return fiber.Map{
		"success": success,
		"message": message,
	}
}
