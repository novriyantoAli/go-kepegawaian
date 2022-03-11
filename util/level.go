package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"
	"github.com/novriyantoAli/go-kepegawaian/model"

	"github.com/casbin/casbin/v2"
	"github.com/golang-jwt/jwt"
)

func LevelProtected(action string, enforcer *casbin.Enforcer) func(c *fiber.Ctx) error {
	return skip.New(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(model.WebResponse{
				Code:   fiber.StatusUnauthorized,
				Status: fiber.ErrUnauthorized.Message,
				Data:   "token or level unauthorized...",
			})
		},

		func(c *fiber.Ctx) bool {
			token, err := verifyToken(c)
			if err != nil {
				return false
			}

			// Setting and checking token and credentials.
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {

				id := string(claims["id"].(string))

				if err = enforcer.LoadPolicy(); err != nil {
					return false
				}

				ok, err := enforcer.Enforce(id, c.Path(), action)
				if err != nil {
					return false
				}

				return ok
			}

			return false
		})
}
