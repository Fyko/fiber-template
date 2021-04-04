package router

import (
	"github.com/fyko/fiber-template/internal/config"
	"github.com/fyko/fiber-template/pkg/handler"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(config config.Config) {
	config.Fiber.Use(cors.New())
	config.Fiber.Use(compress.New())

	api := config.Fiber.Group("/api")
	api.Use("/ping", handler.Ping)
}
