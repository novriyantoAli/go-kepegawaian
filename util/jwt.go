package util

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/novriyantoAli/go-kepegawaian/model"
	"github.com/sirupsen/logrus"

	jwtMiddleware "github.com/gofiber/jwt/v2"
)

func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey:   "jwt", // used in private routes
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	logrus.Warning(err)

	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: fiber.ErrBadRequest.Message,
			Data:   err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(model.WebResponse{
		Code:   fiber.StatusUnauthorized,
		Status: fiber.ErrUnauthorized.Message,
		Data:   err.Error(),
	})
}
