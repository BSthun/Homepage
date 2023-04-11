package enum

type SpotifyEntity string

const (
	SpotifyEntityTrack   SpotifyEntity = "track"
	SpotifyEntityArtist  SpotifyEntity = "artist"
	SpotifyEntityContext SpotifyEntity = "context"
	SpotifyEntityDevice  SpotifyEntity = "device"
	SpotifyEntityRepeat  SpotifyEntity = "repeat"
)
