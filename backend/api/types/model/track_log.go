package model

import (
	"time"
)

type TrackLog struct {
	TrackSession   *TrackSession `gorm:"foreignKey:TrackSessionId" json:"track_session"`
	TrackSessionId *uint64       `gorm:"not null" json:"track_session_id"`
	Tag            *string       `gorm:"type:TEXT; not null" json:"tag"`
	Detail         *string       `gorm:"type:JSON; not null" json:"detail"`
	Error          *string       `gorm:"type:TEXT; null" json:"error"`
	CreatedAt      *time.Time    `gorm:"not null"` // Embedded field
	UpdatedAt      *time.Time    `gorm:"not null"` // Embedded field
}
