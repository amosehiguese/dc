package auth_route

import "github.com/gofiber/fiber/v3"

func authenticate(c fiber.Ctx) error {
	return c.Next()
}
