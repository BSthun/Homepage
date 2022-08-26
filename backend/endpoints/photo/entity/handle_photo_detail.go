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
	var query *payload.SectionPhotoQuery
	if err := c.QueryParser(&query); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}
	s.SetDetail("section-id", query.SectionId)
	s.SetDetail("page-no", query.PageNo)

	// * Query photo detail
	var photoItems []*model.PhotoItem
	if result := mysql.DB.Find(&photoItems, "photo_section_id = ?", query.SectionId).Offset(20 * query.PageNo).Limit(20); result.Error != nil {
		return response.Error(false, "Unable to query photo items", result.Error)
	}

	// * Query photo count
	var photoCount uint64
	if result := mysql.DB.Model(new(model.PhotoItem)).Select("COUNT(*)").First(&photoCount, "photo_section_id = ?", query.SectionId); result.Error != nil {
		return response.Error(false, "Unable to query photo items", result.Error)
	}

	items, _ := value.Iterate(photoItems, func(photoItem *model.PhotoItem) (*payload.PhotoItem, *response.ErrorInstance) {
		splitPath := strings.Split(*photoItem.ImagePath, "/")
		return &payload.PhotoItem{
			Id:            *photoItem.Id,
			Title:         splitPath[len(splitPath)-1],
			ImagePath:     *photoItem.ImagePath,
			ThumbnailPath: *photoItem.ThumbnailPath,
			RawPath:       *photoItem.RawPath,
			Ratio:         *photoItem.Ratio,
			Exif:          photoItem.Exif,
		}, nil
	})

	_ = s.Commit(nil)
	return c.JSON(response.Info(map[string]any{
		"items": items,
	}))
}