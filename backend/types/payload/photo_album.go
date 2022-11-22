package payload

type AlbumList struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	SectionCount uint32 `json:"section_count"`
	ThumbnailUrl string `json:"thumbnail_url"`
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
