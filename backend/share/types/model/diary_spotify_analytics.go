package model

import "share/types/enum"

type DiarySpotifyAnalytics struct {
	Date   *string             `gorm:"type:VARCHAR(255); index:entity,unique; not null" json:"timestamp"` // yyyy-mm-dd
	Entity *enum.SpotifyEntity `gorm:"type:ENUM('track', 'artist', 'context', 'device', 'repeat'); index:entity; not null" json:"entity_type"`
	Value  *string             `gorm:"type:VARCHAR(255); index:entity; not null" json:"entity_id"`
	Count  *uint64             `gorm:"not null" json:"count"`
}
