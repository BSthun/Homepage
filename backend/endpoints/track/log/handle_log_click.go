package log

import (
	"github.com/gofiber/fiber/v2"

	"backend/types/common"
	"backend/types/payload"
	"backend/types/response"
)

func SectionDetailHandler(c *fiber.Ctx) error {
	// * Parse session
	s := c.Locals("session").(*common.Session)
	s.SetTag("track/log/click")

	// * Parse body
	body := new(payload.LogBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	s.SetDetail("ev", body.Event)
	if body.BeginningState != "" {
		s.SetDetail("begin", body.BeginningState)
	}
	s.SetDetail("end", body.EndingState)

	_ = s.Commit(nil)
	return c.JSON(response.Info(map[string]any{
		"accepted": true,
	}))
}