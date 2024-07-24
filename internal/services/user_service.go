package services

import (
	"log"

	"github.com/Adi-111/spotifyDev/internal/config"
	"github.com/Adi-111/spotifyDev/internal/models"
)

// GetAllUsers retrieves all users from the database.
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		log.Printf("Error retrieving users: %v", result.Error)
		return nil, result.Error
	}
	return users, nil
}
