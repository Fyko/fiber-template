package handler

import (
	"github.com/fyko/fiber-template/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(util.CreateReSTResponse(true, "pong"))
}
