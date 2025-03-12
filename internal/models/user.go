package models

import "gorm.io/gorm"

// User represents a user in the system
// It includes fields for ID, Username, Password, and CreatedAt
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string
}
