package photoEntity

import (
	"github.com/gofiber/fiber/v2"

	"backend/loaders/mysql"
	"backend/procedures/photo"
	"backend/types/common"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
)

func SectionDetailHandler(c *fiber.Ctx) error {
	// * Parse session
	s := c.Locals("session").(*common.Session)
	s.SetTag("photo/entity/view-section-detail")

	// * Parse query
	var query *payload.SectionQuery
	if err := c.QueryParser(&query); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}
	s.SetDetail("section-id", query.SectionId)

	// * Query section detail
	var photoSection *model.PhotoSection
	if result := mysql.DB.Preload("PhotoAlbum").First(&photoSection, query.SectionId); result.Error != nil {
		return response.Error(false, "Unable to query album section", result.Error)
	}

	// * Check album permission
	if photoSection.PhotoAlbum.AccessToken != nil {
		if err := photo.CheckAlbumAccess(s, *photoSection.PhotoAlbum.Id, *photoSection.PhotoAlbum.AccessToken, ""); err != nil {
			return err
		}
	}

	// * Generate photo access token
	// TODO: Implement JWT access token

	_ = s.Commit(nil)
	return c.JSON(response.Info(map[string]any{
		"section": &payload.AlbumSection{
			Id:        *photoSection.Id,
			Title:     *photoSection.Title,
			Subtitle:  *photoSection.Subtitle,
			Date:      photoSection.Date,
			UpdatedAt: photoSection.UpdatedAt,
		},
		"album": &payload.AlbumDetail{
			Id:       *photoSection.PhotoAlbum.Id,
			Name:     *photoSection.PhotoAlbum.Name,
			Slug:     *photoSection.PhotoAlbum.Slug,
			Sections: nil,
		},
	}))
}
