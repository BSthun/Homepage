package model

import "time"

type PhotoAlbum struct {
	Id            *uint64    `gorm:"primaryKey" json:"id"`
	Name          *string    `gorm:"type:VARCHAR(255); not null" json:"name"`
	Slug          *string    `gorm:"type:VARCHAR(255); index:slug,unique; not null" json:"slug"`
	AccessToken   *string    `gorm:"type:VARCHAR(255); null" json:"access_token"`
	CoverPhotoUrl *string    `gorm:"type:TEXT" json:"cover_photo_url"`
	SectionCount  *uint64    `gorm:"not null" json:"section_count"`
	PhotoCount    *uint64    `gorm:"not null" json:"photo_count"`
	CreatedAt     *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt     *time.Time `gorm:"not null"` // Embedded field
}
