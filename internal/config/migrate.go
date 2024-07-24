package config

import "github.com/Adi-111/spotifyDev/internal/models"

func MigrateDB() {

	DB.AutoMigrate(&models.User{})

}
