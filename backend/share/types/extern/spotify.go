package extern

type SpotifyPlaybackState struct {
	Device               *SpotifyDevice  `json:"device"`
	ShuffleState         bool            `json:"shuffle_state"`
	RepeatState          string          `json:"repeat_state"`
	Timestamp            int64           `json:"timestamp"`
	Context              *SpotifyContext `json:"context"`
	ProgressMs           int64           `json:"progress_ms"`
	Item                 *SpotifyItem    `json:"item"`
	CurrentlyPlayingType string          `json:"currently_playing_type"`
	IsPlaying            bool            `json:"is_playing"`
}

type SpotifyDevice struct {
	Id               string `json:"id"`
	IsActive         bool   `json:"is_active"`
	IsPrivateSession bool   `json:"is_private_session"`
	IsRestricted     bool   `json:"is_restricted"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	VolumePercent    int    `json:"volume_percent"`
}

type SpotifyExternalUrls struct {
	Spotify string `json:"spotify"`
}

type SpotifyExternalIds struct {
	Isrc string `json:"isrc"`
}

type SpotifyContext struct {
	ExternalUrls *SpotifyExternalUrls `json:"external_urls"`
	Href         string               `json:"href"`
	Type         string               `json:"type"`
	Uri          string               `json:"uri"`
}

type SpotifyItem struct {
	Album        *SpotifyAlbum        `json:"album"`
	Artists      []*SpotifyArtist     `json:"artists"`
	DiscNumber   int                  `json:"disc_number"`
	DurationMs   int                  `json:"duration_ms"`
	Explicit     bool                 `json:"explicit"`
	ExternalIds  *SpotifyExternalIds  `json:"external_ids"`
	ExternalUrls *SpotifyExternalUrls `json:"external_urls"`
	Href         string               `json:"href"`
	Id           string               `json:"id"`
	IsLocal      bool                 `json:"is_local"`
	Name         string               `json:"name"`
	Popularity   int                  `json:"popularity"`
	PreviewUrl   string               `json:"preview_url"`
	TrackNumber  int                  `json:"track_number"`
	Type         string               `json:"type"`
	Uri          string               `json:"uri"`
}

type SpotifyAlbum struct {
	AlbumType            string               `json:"album_type"`
	Artists              []*SpotifyArtist     `json:"artists"`
	AvailableMarkets     []string             `json:"available_markets"`
	ExternalUrls         *SpotifyExternalUrls `json:"external_urls"`
	Href                 string               `json:"href"`
	Id                   string               `json:"id"`
	Images               []*SpotifyImage      `json:"images"`
	Name                 string               `json:"name"`
	ReleaseDate          string               `json:"release_date"`
	ReleaseDatePrecision string               `json:"release_date_precision"`
	TotalTracks          int                  `json:"total_tracks"`
	Type                 string               `json:"type"`
	Uri                  string               `json:"uri"`
}

type SpotifyArtist struct {
	ExternalUrls *SpotifyExternalUrls `json:"external_urls"`
	Href         string               `json:"href"`
	Id           string               `json:"id"`
	Name         string               `json:"name"`
	Type         string               `json:"type"`
	Uri          string               `json:"uri"`
}

type SpotifyImage struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Url    string `json:"url"`
}
