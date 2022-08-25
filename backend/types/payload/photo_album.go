package payload

type AlbumDetail struct {
	Id       uint64          `json:"id"`
	Name     string          `json:"name"`
	Slug     string          `json:"slug"`
	Sections []*AlbumSection `json:"sections"`
}
