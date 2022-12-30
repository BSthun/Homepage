package diary

import (
	"github.com/gofiber/fiber/v2"

	"backend/types/common"
)

func OverviewHandler(c *fiber.Ctx) error {
	// * Parse session
	s := c.Locals("session").(*common.Session)
	s.SetTag("diary/overview")

	return nil
}
