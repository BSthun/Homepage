package model

import "time"

type PhotoItemPerson struct {
	PhotoItem     *PhotoItem   `gorm:"foreignKey:PhotoItemId" json:"photo_item"`
	PhotoItemId   *uint64      `gorm:"primaryKey" json:"photo_item_id"`
	PhotoPerson   *PhotoPerson `gorm:"foreignKey:PhotoPersonId" json:"photo_person"`
	PhotoPersonId *uint64      `gorm:"primaryKey" json:"photo_person_id"`
	CreatedAt     *time.Time   `gorm:"not null"` // Embedded field
	UpdatedAt     *time.Time   `gorm:"not null"` // Embedded field
}
