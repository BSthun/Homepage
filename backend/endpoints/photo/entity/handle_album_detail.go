package photoEntity

import (
	"github.com/gofiber/fiber/v2"

	"backend/loaders/mysql"
	"backend/procedures/photo"
	"backend/types/common"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/value"
)

func AlbumDetailHandler(c *fiber.Ctx) error {
	// * Parse session
	s := c.Locals("session").(*common.Session)
	s.SetTag("photo/entity/view-album-detail")

	// * Parse query
	var query *payload.AlbumQuery
	if err := c.QueryParser(&query); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}
	s.SetDetail("album-slug", query.AlbumSlug)

	// * Query album detail
	var photoAlbum *model.PhotoAlbum
	if result := mysql.DB.First(&photoAlbum, "slug = ?", query.AlbumSlug); result.Error != nil {
		return response.Error(false, "Unable to query album detail", result.Error)
	}

	// * Check album permission
	if photoAlbum.AccessToken != nil {
		if err := photo.CheckAlbumAccess(s, *photoAlbum.Id, *photoAlbum.AccessToken, query.Token); err != nil {
			return err
		}
	}

	// * Query album sections
	var albumSections []*payload.AlbumSection
	if result := mysql.DB.Find(&albumSections, "photo_album_id = ?", photoAlbum.Id); result.Error != nil {
		return response.Error(true, "Unable to query list of album sections")
	}

	sections, _ := value.Iterate(albumSections, func(albumSection *payload.AlbumSection) (*payload.AlbumSection, *response.ErrorInstance) {
		return &payload.AlbumSection{
			Id:        albumSection.Id,
			Title:     albumSection.Title,
			Subtitle:  albumSection.Subtitle,
			Date:      albumSection.Date,
			UpdatedAt: albumSection.UpdatedAt,
		}, nil
	})

	_ = s.Commit(nil)
	return c.JSON(response.Info(map[string]any{
		"album": &payload.AlbumDetail{
			Id:       *photoAlbum.Id,
			Name:     *photoAlbum.Name,
			Slug:     *photoAlbum.Slug,
			Sections: sections,
		},
	}))
}
