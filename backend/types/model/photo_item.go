package model

import (
	"time"

	"backend/types/presentation"
)

type PhotoItem struct {
	Id             *uint64            `gorm:"primaryKey" json:"id"`
	PhotoSection   *PhotoSection      `gorm:"foreignKey:PhotoSectionId" json:"photo_section"`
	PhotoSectionId *uint64            `gorm:"not null" json:"photo_section_id"`
	Shooter        *PhotoPerson       `gorm:"foreignKey:ShooterId" json:"shooter"`
	ShooterId      *uint64            `gorm:"<-:create; null" json:"shooter_id"`
	Root           *string            `gorm:"type:TEXT; not null" json:"root"`
	ImagePath      *string            `gorm:"type:TEXT; not null" json:"image_path"`
	ThumbnailPath  *string            `gorm:"type:TEXT; not null" json:"thumbnail_path"`
	RawPath        *string            `gorm:"type:TEXT; not null" json:"raw_path"`
	Exif           *presentation.Exif `gorm:"type:TEXT; not null" json:"exif"`
	Ratio          *float32           `gorm:"type:FLOAT; not null" json:"ratio"`
	CreatedAt      *time.Time         `gorm:"not null"` // Embedded field
	UpdatedAt      *time.Time         `gorm:"not null"` // Embedded field
}
