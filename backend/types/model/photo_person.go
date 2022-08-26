package model

import "time"

type PhotoPerson struct {
	Id         *uint64    `gorm:"primaryKey" json:"id"`
	Name       *string    `gorm:"type:VARCHAR(255); not null" json:"name"`
	AvatarPath *string    `gorm:"type:TEXT; not null" json:"avatar_path"`
	CreatedAt  *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt  *time.Time `gorm:"not null"` // Embedded field
}
