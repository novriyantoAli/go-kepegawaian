package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novriyantoAli/go-kepegawaian/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
