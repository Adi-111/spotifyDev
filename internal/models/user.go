package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Username    string `gorm:"unique;not null"`
	Email       string `gorm:"unique;not null"`
	Password    string
	Country     string `gorm:"not null"`
	SpotifyID   string `gorm:"uniqueIndex;not null"`
	Href        string `gorm:"not null"`
	Type        string `gorm:"not null"`
	DisplayName string `gorm:"not null"`
}
