package photoEntity

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"backend/loaders/mysql"
	"backend/types/common"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/value"
)

func PhotoListHandler(c *fiber.Ctx) error {
	// * Parse session
	s := c.Locals("session").(*common.Session)
	s.SetTag("photo/entity/view-photo-list")

	// * Parse query
	query := new(payload.SectionPhotoQuery)
	if err := c.QueryParser(query); err != nil {
		return response.Error(false, "Unable to parse query", err)
	}
	s.SetDetail("section-id", query.SectionId)
	s.SetDetail("page-no", query.PageNo)

	// * Query photo detail
	var photoItems []*model.PhotoItem
	if result := mysql.DB.Offset(20*query.PageNo).Limit(20).Preload("PhotoSection").Find(&photoItems, "photo_section_id = ?", query.SectionId); result.Error != nil {
		return response.Error(false, "Unable to query photo items", result.Error)
	}

	items, _ := value.Iterate(photoItems, func(photoItem *model.PhotoItem) (*payload.PhotoItem, *response.ErrorInstance) {
		splitPath := strings.Split(*photoItem.ImagePath, "/")
		return &payload.PhotoItem{
			Id:            *photoItem.Id,
			Title:         splitPath[len(splitPath)-1],
			Root:          *photoItem.PhotoSection.Path,
			ImagePath:     *photoItem.ImagePath,
			ThumbnailPath: *photoItem.ThumbnailPath,
			RawPath:       *photoItem.RawPath,
			Exif:          photoItem.Exif,
		}, nil
	})

	_ = s.Commit(nil)
	return c.JSON(response.Info(map[string]any{
		"items": items,
	}))
}
