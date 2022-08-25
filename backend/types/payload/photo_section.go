package payload

import "time"

type AlbumSection struct {
	Id         uint64     `json:"id"`
	Title      string     `json:"title"`
	Subtitle   string     `json:"subtitle"`
	Date       *time.Time `json:"date"`
	PhotoCount *uint32    `json:"photo_count"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

type SectionQuery struct {
	SectionId uint64 `json:"section_id" query:"section_id"`
}

type SectionPhotoQuery struct {
	SectionId uint64 `json:"section_id" query:"section_id"`
	PageNo    int    `json:"page_no" query:"page_no"`
}
