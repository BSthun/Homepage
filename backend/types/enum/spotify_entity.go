package enum

type SpotifyEntity string

const (
	SpotifyEntityTrack    SpotifyEntity = "track"
	SpotifyEntityArtist   SpotifyEntity = "artist"
	SpotifyEntityPlaylist SpotifyEntity = "playlist"
	SpotifyEntityDevice   SpotifyEntity = "device"
	SpotifyEntityRepeat   SpotifyEntity = "repeat"
)
