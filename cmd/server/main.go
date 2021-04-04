//go:generate go run github.com/prisma/prisma-client-go generate
package main

import (
	"fmt"
	"os"

	"github.com/fyko/fiber-template/pkg/errors"
	mylog "github.com/fyko/fiber-template/pkg/logger"
	"github.com/fyko/fiber-template/pkg/router"
	"github.com/fyko/fiber-template/prisma/db"

	"github.com/fyko/fiber-template/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var port = os.Getenv("PORT")
var log = mylog.CreateLogger()

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Errorf("Internal server error: %v", err)
			return c.Status(500).JSON(errors.CreateReSTError("Internal Server Error"))
		},
		Immutable: true,
	})
	app.Use(logger.New(logger.Config{
		TimeZone: "America/Denver",
	}))

	db := db.NewClient()
	if err := db.Prisma.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}

	config := config.Config{
		Fiber:  app,
		Logger: log,
		Prisma: db,
	}

	router.SetupRoutes(config)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
