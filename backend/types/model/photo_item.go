package model

import "time"

type PhotoItem struct {
	Id             *uint64       `gorm:"primaryKey" json:"id"`
	PhotoSection   *PhotoSection `gorm:"foreignKey:PhotoSectionId" json:"photo_section"`
	PhotoSectionId *uint64       `gorm:"not null" json:"photo_section_id"`
	ImagePath      *string       `gorm:"type:TEXT; not null" json:"image_path"`
	ThumbnailPath  *string       `gorm:"type:TEXT; not null" json:"thumbnail_path"`
	RawPath        *string       `gorm:"type:TEXT; not null" json:"raw_path"`
	PhotoCamera    *PhotoCamera  `gorm:"foreignKey:PhotoCameraId" json:"photo_camera"`
	PhotoCameraId  *uint64       `gorm:"not null" json:"photo_camera_id"`
	Shooter        *PhotoPerson  `gorm:"foreignKey:ShooterId" json:"shooter"`
	ShooterId      *uint64       `gorm:"<-:create; null" json:"shooter_id"`
	Exif           *string       `gorm:"type:TEXT; not null" json:"exif"`
	CreatedAt      *time.Time    `gorm:"not null"` // Embedded field
	UpdatedAt      *time.Time    `gorm:"not null"` // Embedded field
}
