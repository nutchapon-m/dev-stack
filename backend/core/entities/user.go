package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           string `gorm:"primaryKey;type:varchar(191)"`
	Username     string
	HashPassword string
	Name         string
	Email        string
	CreatedBy    string
	CreatedAt    string `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP()"`
	UpdatedBy    string
	UpdatedAt    string `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP()"`
	Status       int    `gorm:"default:1"`
}
