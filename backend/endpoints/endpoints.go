package endpoints

import (
	"github.com/gofiber/fiber/v2"

	"backend/endpoints/account/state"
	photoEntity "backend/endpoints/photo/entity"
	trackLog "backend/endpoints/track/log"
	"backend/loaders/fiber/middlewares"
)

func Init(router fiber.Router) {
	// * Account
	account := router.Group("/account", middlewares.Session)
	account.Get("/state", accountState.StateHandler)

	// * Photo
	photo := router.Group("/photo", middlewares.Session)
	photo.Get("/entity/album/list", photoEntity.AlbumListHandler)
	photo.Get("/entity/album/detail", photoEntity.AlbumDetailHandler)
	photo.Get("/entity/section/detail", photoEntity.SectionDetailHandler)
	photo.Get("/entity/photo/list", photoEntity.PhotoListHandler)

	// * Tracking
	tracking := router.Group("/track", middlewares.Session)
	tracking.Put("/log/click", trackLog.SectionDetailHandler)
	tracking.Put("/sig/err", trackLog.SigErrorHandler)
}
