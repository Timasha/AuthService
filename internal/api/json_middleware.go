package api

import (
	"github.com/gofiber/fiber/v2"
)

func (a *Auth) GetJsonMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !c.Is("json") {
			return fiber.ErrUnsupportedMediaType
		}
		return c.Next()
	}
}
