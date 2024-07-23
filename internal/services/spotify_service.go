package services

import (
	"golang.org/x/oauth2"
)

var (
	SpotifyAuth = oauth2.Config{
		ClientID:     "b92abedbb5e14b65a3432a8ace28898b",
		ClientSecret: "808ada812a26418684cf9d2c808aff8b",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.spotify.com/authorize",
			TokenURL: "https://accounts.spotify.com/api/token",
		},
		RedirectURL: "http://localhost:8080/callback",
		Scopes:      []string{"user-read-private", "user-read-email"},
	}
)
