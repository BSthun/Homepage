package model

type PhotoSessionAlbum struct {
	TrackSessionId *uint64       `gorm:"primaryKey" json:"track_session_id"`
	TrackSession   *TrackSession `gorm:"foreignKey:TrackSessionId" json:"track_session"`
	PhotoAlbumId   *uint64       `gorm:"primaryKey" json:"photo_album_id"`
	PhotoAlbum     *PhotoAlbum   `gorm:"foreignKey:PhotoAlbumId" json:"photo_album"`
}
