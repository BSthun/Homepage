package payload

import "backend/types/presentation"

type PhotoItem struct {
	Id            uint64             `json:"id"`
	Title         string             `json:"title"`
	Root          string             `json:"root"`
	ImagePath     string             `json:"image_path"`
	ThumbnailPath string             `json:"thumbnail_path"`
	RawPath       string             `json:"raw_path"`
	Ratio         float32            `json:"ratio"`
	Exif          *presentation.Exif `json:"exif"`
}

type PhotoQuery struct {
	PhotoId uint64 `json:"photo_id" query:"photo_id"`
}
