package photoEntity

import (
	"github.com/gofiber/fiber/v2"

	"backend/loaders/mysql"
	"backend/types/common"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/value"
)

func AlbumListHandler(c *fiber.Ctx) error {
	// * Parse session
	s := c.Locals("session").(*common.Session)
	s.SetTag("photo/entity/view-album-list")

	// * Query photo detail
	var photoAlbums []*model.PhotoAlbum
	if result := mysql.DB.Preload("CoverPhoto.PhotoSection").Find(&photoAlbums); result.Error != nil {
		return response.Error(false, "Unable to query photo albums", result.Error)
	}

	albums, _ := value.Iterate(photoAlbums, func(photoItem *model.PhotoAlbum) (*payload.AlbumList, *response.ErrorInstance) {
		return &payload.AlbumList{
			Id:           *photoItem.Id,
			Name:         *photoItem.Name,
			Slug:         *photoItem.Slug,
			SectionCount: *photoItem.SectionCount,
			PhotoCount:   *photoItem.PhotoCount,
			ThumbnailUrl: *photoItem.CoverPhoto.PhotoSection.Path + *photoItem.CoverPhoto.ThumbnailPath,
			UpdatedAt:    *photoItem.UpdatedAt,
		}, nil
	})

	_ = s.Commit(nil)
	return c.JSON(response.Info(map[string]any{
		"albums": albums,
	}))
}
