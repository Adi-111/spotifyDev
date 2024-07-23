package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Adi-111/spotifyDev/internal/services"
	"golang.org/x/oauth2"
)

func SpotifyLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("initiating spotify login")
	url := services.SpotifyAuth.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("url genrated : %s", url)

	http.Redirect(w, r, url, http.StatusSeeOther)
}
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("waiting for callback")
	code := r.URL.Query().Get("code")
	token, err := services.SpotifyAuth.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to authorize token", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "Access Token: %s", token.AccessToken)
	// Get Spotify user information
	user, err := services.GetSpotifyUserInfo(token.AccessToken)
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}

	// Display user information
	fmt.Fprintf(w, "User Info:\n")
	fmt.Fprintf(w, "Display Name: %s\n", user.DisplayName)
	fmt.Fprintf(w, "ID: %s\n", user.ID)
	fmt.Fprintf(w, "Email: %s\n", user.Email)
}
