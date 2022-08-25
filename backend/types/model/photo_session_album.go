package model

type PhotoSessionAlbum struct {
	TrackSession   *TrackSession `gorm:"foreignKey:TrackSessionId" json:"track_session"`
	TrackSessionId *uint64       `gorm:"primaryKey" json:"track_session_id"`
	PhotoAlbum     *PhotoAlbum   `gorm:"foreignKey:PhotoAlbumId" json:"photo_album"`
	PhotoAlbumId   *uint64       `gorm:"primaryKey" json:"photo_album_id"`
}
