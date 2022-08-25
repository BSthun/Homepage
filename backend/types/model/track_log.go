package model

import (
	"time"

	"backend/types/enum"
)

type TrackLog struct {
	TrackSession   *TrackSession `gorm:"foreignKey:TrackSessionId" json:"track_session"`
	TrackSessionId *uint64       `gorm:"primaryKey" json:"track_session_id"`
	Tag            *enum.Tag     `gorm:"type:ENUM('photo'); not null" json:"tag"`
	Detail         *string       `gorm:"type:TEXT; not null" json:"detail"`
	CreatedAt      *time.Time    `gorm:"not null"` // Embedded field
	UpdatedAt      *time.Time    `gorm:"not null"` // Embedded field
}
