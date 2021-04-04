package config

import (
	"github.com/fyko/fiber-template/prisma/db"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Config struct {
	Fiber  *fiber.App
	Prisma *db.PrismaClient
	Logger *zap.SugaredLogger
}
