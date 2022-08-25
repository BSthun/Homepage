package payload

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
