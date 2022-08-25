package accountState

import (
	"github.com/gofiber/fiber/v2"

	"backend/types/response"
)

func StateHandler(c *fiber.Ctx) error {
	return c.JSON(response.Info(map[string]any{
		"state": "ok",
	}))
}
