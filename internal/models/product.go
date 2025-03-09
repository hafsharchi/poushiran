package models

import (
	"time"

	"gorm.io/gorm"
)

// Product represents a product in the system
type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" binding:"required" gorm:"size:100;not null"`
	Description string         `json:"description" binding:"required" gorm:"size:500"`
	Price       float64        `json:"price" binding:"required,gt=0"`
	Category    string         `json:"category" binding:"required" gorm:"size:50"`
	InStock     bool           `json:"inStock" gorm:"default:true"`
	CreatedAt   time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
