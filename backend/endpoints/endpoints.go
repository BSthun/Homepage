package endpoints

import (
	"github.com/gofiber/fiber/v2"

	"backend/endpoints/account/state"
	photoEntity "backend/endpoints/photo/entity"
	"backend/loaders/fiber/middlewares"
)

func Init(router fiber.Router) {
	// * Account
	account := router.Group("/account", middlewares.Session)
	account.Get("/state", accountState.StateHandler)

	// * Photo
	photo := router.Group("/photo")
	photo.Get("/entity/album/detail", photoEntity.AlbumDetailHandler)
	photo.Get("/entity/section/detail", photoEntity.SectionDetailHandler)
	photo.Get("/entity/photo/list", photoEntity.PhotoListHandler)
}
