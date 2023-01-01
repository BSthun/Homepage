package trackLog

import (
	"github.com/gofiber/fiber/v2"

	"backend/types/common"
	"backend/types/payload"
	"backend/types/response"
)

func SigErrorHandler(c *fiber.Ctx) error {
	// * Parse session
	s := c.Locals("session").(*common.Session)
	s.SetTag("track/sig/error")

	// * Parse body
	body := new(payload.SigBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	s.SetDetail("t", body.Title)

	_ = s.Commit(nil)
	return c.JSON(response.Info(map[string]any{
		"accepted": true,
	}))
}
