package model

import "time"

type TrackSession struct {
	Id         *uint64    `gorm:"primaryKey" json:"id"`
	Token      *string    `gorm:"type:TEXT; not null" json:"token"`
	UserAgent  *string    `gorm:"type:TEXT; not null" json:"user_agent"`
	IpAddress  *string    `gorm:"type:TEXT; not null" json:"ip_address"`
	Attributes *string    `gorm:"type:JSON; not null" json:"attributes"`
	CreatedAt  *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt  *time.Time `gorm:"not null"` // Embedded field
}
