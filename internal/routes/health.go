package routes

import (
	"github.com/gabrielluciano/rinha-backend-2024-q1-go/internal/util"
	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	if !util.Warm {
		return c.Status(500).Send(nil)
	}
	return c.Status(200).Send(nil)
}
