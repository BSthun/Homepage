package fiber

import (
	"github.com/gofiber/fiber/v2"

	"backend/types/response"
)

func NotFoundHandler(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
		Success: false,
		Error:   "404_NOT_FOUND",
	})
}
