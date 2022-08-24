package model

import "time"

type PhotoCamera struct {
	Id          *uint64    `gorm:"primaryKey" json:"id"`
	CameraMaker *string    `gorm:"type:VARCHAR(255); not null" json:"camera_maker"`
	CameraModel *string    `gorm:"type:VARCHAR(255); not null" json:"camera_name"`
	LensModel   *string    `gorm:"type:VARCHAR(255); not null" json:"lens_model"`
	CreatedAt   *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt   *time.Time `gorm:"not null"` // Embedded field
}
