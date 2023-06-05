package payload

import (
	"time"

	"backend/types/model"
)

type ExtendedPhotoSection struct {
	model.PhotoSection
	PhotoCount   *uint64 `json:"photo_count"`
	ThumbnailUrl *string `json:"thumbnail_url"`
}

type AlbumSection struct {
	Id           uint64    `json:"id"`
	Title        string    `json:"title"`
	Subtitle     string    `json:"subtitle"`
	Date         time.Time `json:"date"`
	PhotoCount   uint64    `json:"photo_count"`
	ThumbnailUrl string    `json:"thumbnail_url"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type SectionQuery struct {
	SectionId uint64 `json:"section_id" query:"section_id"`
	Token     string `json:"token" query:"token"`
}

type SectionPhotoQuery struct {
	SectionId uint64 `json:"section_id" query:"section_id"`
	PageNo    int    `json:"page_no" query:"page_no"`
}
