package middlewares

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"backend/loaders/mysql"
	"backend/types/common"
	"backend/types/model"
	"backend/types/response"
	"backend/utils/text"
	"backend/utils/value"
)

var Session = func() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var session *model.TrackSession
		var newSession bool
		if c.Cookies("session_id") == "" {
			newSession = true
		} else {
			if result := mysql.DB.First(&session, c.Cookies("session_id")); result.Error != nil {
				newSession = true
			} else if *session.Token != c.Cookies("session_token") {
				newSession = true
			}
		}

		if newSession {
			session = &model.TrackSession{
				Id:        nil,
				Token:     text.Random(text.RandomSet.MixedAlphaNum, 16),
				UserAgent: value.Ptr(c.Get("User-Agent")),
				IpAddress: value.Ptr(c.Get("X-Real-IP")),
				CreatedAt: nil,
				UpdatedAt: nil,
			}
			if result := mysql.DB.Create(session); result.Error != nil {
				return response.Error(true, "Failed to create session", result.Error)
			}

			c.Cookie(&fiber.Cookie{
				Name:        "session_id",
				Value:       strconv.FormatUint(*session.Id, 10),
				Path:        "/",
				Domain:      "",
				MaxAge:      0,
				Expires:     time.Now().Add(7 * 24 * time.Hour),
				Secure:      false,
				HTTPOnly:    false,
				SameSite:    "",
				SessionOnly: false,
			})
			c.Cookie(&fiber.Cookie{
				Name:        "session_token",
				Value:       *session.Token,
				Path:        "/",
				Domain:      "",
				MaxAge:      0,
				Expires:     time.Now().Add(7 * 24 * time.Hour),
				Secure:      false,
				HTTPOnly:    false,
				SameSite:    "",
				SessionOnly: false,
			})
		}

		c.Locals("session", &common.Session{
			Id: *session.Id,
		})

		return c.Next()
	}
}()
