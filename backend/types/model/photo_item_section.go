package model

import "time"

type PhotoItemSection struct {
	PhotoItem      *PhotoItem    `gorm:"foreignKey:PhotoItemId" json:"photo_item"`
	PhotoItemId    *uint64       `gorm:"primaryKey" json:"photo_item_id"`
	PhotoSection   *PhotoSection `gorm:"foreignKey:PhotoSectionId" json:"photo_section"`
	PhotoSectionId *uint64       `gorm:"primaryKey" json:"photo_section_id"`
	SectionId      *uint64       `gorm:"primaryKey" json:"section_id"`
	CreatedAt      *time.Time    `gorm:"not null"` // Embedded field
	UpdatedAt      *time.Time    `gorm:"not null"` // Embedded field
}
