package hub

import (
	"strings"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"

	"backend/types/enum"
)

var Hub *hubModel

func Init() {
	Hub = &hubModel{
		SpotifyAccessToken:  "",
		SpotifyTokenExpired: time.Now(),
	}

	if err := spotifyToken(); err != nil {
		logrus.Warn("Unable to get Spotify token: ", err)
	}

	s := gocron.NewScheduler(time.UTC)
	_, _ = s.Every(1 * time.Minute).Do(func() {
		state, err := spotifyPlayback()
		if err != nil {
			logrus.Warn("Unable to get Spotify playback state: ", err)
			return
		}

		// * Check if playback is active
		if !state.IsPlaying {
			return
		}

		// * Check if playback is a track
		if state.Item.Type != "track" {
			return
		}

		// * Check if track is registered
		spotifyRecord(enum.SpotifyEntityTrack, state.Item.Id)
		for _, artist := range state.Item.Artists {
			spotifyRecord(enum.SpotifyEntityArtist, artist.Id)
		}
		if state.Context.Type == "playlist" {
			spotifyRecord(enum.SpotifyEntityPlaylist, strings.Split(state.Context.Uri, ":")[2])
		}
		spotifyRecord(enum.SpotifyEntityDevice, state.Device.Name)
		spotifyRecord(enum.SpotifyEntityRepeat, state.RepeatState)
	})
	s.StartAsync()
}
