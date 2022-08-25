package photoEntity

import (
	"github.com/gofiber/fiber/v2"

	"backend/types/common"
	"backend/types/enum"
)

func AlbumDetailHandler(c *fiber.Ctx) error {
	// * Log session
	s := c.Locals("session").(*common.Session)
	_ = s.Log(enum.TagPhoto, map[string]any{})

	return nil
}
