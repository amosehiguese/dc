package user

import "github.com/gofiber/fiber/v3"

func UserRoutes(v1 fiber.Router) {
	users := v1.Group("/users")
	users.Get("/", func(c fiber.Ctx) error {
		return c.JSON(map[string]any{
			"msg": "Got to user routes",
		})
	})
}
