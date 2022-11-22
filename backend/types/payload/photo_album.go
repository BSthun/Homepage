package payload

import "time"

type AlbumList struct {
	Id           uint64    `json:"id"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"`
	ThumbnailUrl string    `json:"thumbnail_url"`
	SectionCount uint64    `json:"section_count"`
	PhotoCount   uint64    `json:"photo_count"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AlbumDetail struct {
	Id       uint64          `json:"id"`
	Name     string          `json:"name"`
	Slug     string          `json:"slug"`
	Sections []*AlbumSection `json:"sections"`
}

type AlbumQuery struct {
	AlbumSlug string `json:"album_slug" query:"album_slug"`
	Token     string `json:"token" query:"token"`
}
