package handlers

import (
	"context"
	"encoding/json"
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

	fmt.Println("\nwaiting for callback")

	code := r.URL.Query().Get("code")
	token, err := services.SpotifyAuth.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to authorize token", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "Access Token: %s", token.AccessToken)
	// Get Spotify user information
	// Fetch the user's profile information from Spotify
	client := services.SpotifyAuth.Client(context.Background(), token)
	resp, err := client.Get("https://api.spotify.com/v1/me")
	if err != nil {
		http.Error(w, "Failed to get user info from Spotify", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	profile, err := services.FetchProfile(token.AccessToken)
	if err != nil {
		http.Error(w, "Failed to fetch profile: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if the user exists, and create if not
	user, err := services.FindOrCreateUser(profile)

	if err != nil {
		http.Error(w, "Failed to find or create user", http.StatusInternalServerError)
		return
	}

	// Return the access token or any other necessary information
	fmt.Println("login succesfull")
	fmt.Fprintf(w, "Access Token: %s\n", token.AccessToken)
	fmt.Fprintf(w, "User: %+v\n", user)
	fmt.Fprintf(w, "Access Token: %s\n", token.AccessToken)
	fmt.Fprintf(w, "Profile ID: %s\n", profile.ID)
	fmt.Fprintf(w, "Email: %s\n", profile.Email)

}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// FetchProfile fetches the user's Spotify profile using the access token
