package hub

import (
	"time"

	"github.com/sirupsen/logrus"
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
}
