package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/Adi-111/spotifyDev/internal/services"
	"golang.org/x/oauth2"
)

func SpotifyLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("initiating spotify login")
	url := services.SpotifyAuth.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("url generated : %s", url)

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nwaiting for callback")

	code := r.URL.Query().Get("code")
	token, err := services.SpotifyAuth.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to authorize token", http.StatusInternalServerError)
		return
	}

	// Fetch the user's profile information from Spotify

	profile, err := services.FetchProfile(token.AccessToken)
	if err != nil {
		http.Error(w, "Failed to fetch profile: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if the user exists, and create if not
	services.FindOrCreateUser(profile)

	// Encode the token and profile as a JSON object
	data := map[string]interface{}{
		"token":   token.AccessToken,
		"profile": profile,
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		return
	}
	dataBody := url.QueryEscape(string(dataBytes))

	// Redirect to success page with dataBody
	redirectURL := fmt.Sprintf("http://localhost:8080/success?dataBody=%s", dataBody)
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
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

func SuccessHandler(w http.ResponseWriter, r *http.Request) {
	dataBody := r.URL.Query().Get("dataBody")
	if dataBody == "" {
		http.Error(w, "No data provided", http.StatusBadRequest)
		return
	}

	dataBytes, err := url.QueryUnescape(dataBody)
	if err != nil {
		http.Error(w, "Failed to decode data", http.StatusInternalServerError)
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(dataBytes), &data); err != nil {
		http.Error(w, "Failed to parse data", http.StatusInternalServerError)
		return
	}

	token := data["token"].(string)
	profile := data["profile"].(map[string]interface{})

	// For debugging, print the current working directory
	cwd, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		http.Error(w, "Failed to get current working directory", http.StatusInternalServerError)
		log.Printf("Error getting current working directory: %v", err)
		return
	}
	log.Printf("Current working directory: %s", cwd)

	// Path to the success.html template
	templatePath := filepath.Join(cwd, "../static", "success.html")
	log.Printf("Template path: %s", templatePath)

	// Parse the template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	// Inject the token and profile into the template and serve it
	templateData := map[string]interface{}{
		"Token":   token,
		"Profile": profile,
	}

	// Execute the template with the token and profile
	err = tmpl.Execute(w, templateData)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}

	log.Println("Template rendered successfully")
}
