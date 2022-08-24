package model

import "time"

type TrackSession struct {
	Id        *uint64    `gorm:"primaryKey" json:"id"`
	UserAgent *string    `gorm:"type:VARCHAR(255); not null" json:"user_agent"`
	IpAddress *string    `gorm:"type:VARCHAR(255); not null" json:"ip_address"`
	CreatedAt *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt *time.Time `gorm:"not null"` // Embedded field
}
