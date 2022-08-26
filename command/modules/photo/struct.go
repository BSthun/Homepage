package photo

import "time"

type PhotoItem struct {
	PhotoSectionId uint64    `gorm:"not null" json:"photo_section_id"`
	Ratio          float32   `gorm:"type:FLOAT; not null" json:"ratio"`
	Root           string    `gorm:"type:TEXT; not null" json:"root"`
	ImagePath      string    `gorm:"type:TEXT; not null" json:"image_path"`
	ThumbnailPath  string    `gorm:"type:TEXT; not null" json:"thumbnail_path"`
	RawPath        string    `gorm:"type:TEXT; not null" json:"raw_path"`
	PhotoCameraId  uint64    `gorm:"not null" json:"photo_camera_id"`
	ShooterId      uint64    `gorm:"<-:create; null" json:"shooter_id"`
	Exif           *Exif     `gorm:"type:TEXT; not null" json:"exif"`
	CreatedAt      time.Time `gorm:"not null"` // Embedded field
	UpdatedAt      time.Time `gorm:"not null"` // Embedded field
}

type ExifRaw struct {
	ApertureValue         []string `json:"ApertureValue"`
	BrightnessValue       []string `json:"BrightnessValue"`
	Artist                string   `json:"Artist"`
	Copyright             string   `json:"Copyright"`
	DateTimeOriginal      string   `json:"DateTimeOriginal"`
	ExposureTime          []string `json:"ExposureTime"`
	FocalLength           []string `json:"FocalLength"`
	FocalLengthIn35mmFilm []int    `json:"FocalLengthIn35mmFilm"`
	IsoSpeedRatings       []int    `json:"ISOSpeedRatings"`
	LensModel             string   `json:"LensModel"`
	Make                  string   `json:"Make"`
	Model                 string   `json:"Model"`
	PixelXDimension       []int    `json:"PixelXDimension"`
	PixelYDimension       []int    `json:"PixelYDimension"`
	Software              string   `json:"Software"`
	WhiteBalance          []int    `json:"WhiteBalance"`
}

type Exif struct {
	Aperture        string    `json:"aperture"`
	Brightness      string    `json:"brightness"`
	Credit          string    `json:"credit"`
	Timestamp       time.Time `json:"timestamp"`
	ExposureTime    string    `json:"exposure_time"`
	FocalLength     string    `json:"focal_length"`
	FocalLength35mm string    `json:"focal_length_35mm"`
	IsoSpeed        string    `json:"iso_speed"`
	LensModel       string    `json:"lens_model"`
	CameraMaker     string    `json:"camera_maker"`
	CameraModel     string    `json:"camera_model"`
}