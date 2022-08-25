package payload

import "time"

type AlbumSection struct {
	Id        uint64     `json:"id"`
	Title     string     `json:"title"`
	Subtitle  string     `json:"subtitle"`
	Date      *time.Time `json:"date"`
	UpdatedAt *time.Time `json:"updated_at"`
}
