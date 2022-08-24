package model

type TrackEvent struct {
	TrackSession   *TrackSession `gorm:"foreignKey:TrackSessionId" json:"track_session"`
	TrackSessionId *uint64       `gorm:"primaryKey" json:"track_session_id"`
}
