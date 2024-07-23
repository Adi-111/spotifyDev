package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SpotifyUser struct {
	DisplayName string `json:"display_name"`
	ID          string `json:"id"`
	Email       string `json:"email"`
}

func GetSpotifyUserInfo(token string) (*SpotifyUser, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user SpotifyUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return &user, nil
}
