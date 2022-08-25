package endpoints

import (
	"github.com/gofiber/fiber/v2"

	"backend/endpoints/account/state"
)

func Init(router fiber.Router) {
	// * Account
	account := router.Group("/account")
	account.Get("/state", accountState.StateHandler)
}
