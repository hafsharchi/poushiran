package database

import (
	"log"

	"github.com/poushiran/config"
	"github.com/poushiran/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Initialize sets up the database connection and runs migrations
func Initialize(config *config.Config) error {
	var err error

	// Connect to SQLite database
	DB, err = gorm.Open(sqlite.Open(config.Database.Path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}

	log.Println("Connected to database")

	// Run migrations
	err = runMigrations()
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
		return err
	}

	log.Println("Database migrations completed")
	return nil
}

// runMigrations runs database migrations
func runMigrations() error {
	// Auto migrate the models
	err := DB.AutoMigrate(&models.Product{}, &models.User{})
	if err != nil {
		return err
	}

	// Seed initial data if the database is empty
	var count int64
	DB.Model(&models.Product{}).Count(&count)
	if count == 0 {
		seedData()
	}

	return nil
}

// seedData seeds the database with initial data
func seedData() {
	products := []models.Product{
		{
			Name:        "Laptop",
			Description: "High-performance laptop with 16GB RAM",
			Price:       999.99,
			Category:    "Electronics",
			InStock:     true,
		},
		{
			Name:        "Smartphone",
			Description: "Latest smartphone with 128GB storage",
			Price:       699.99,
			Category:    "Electronics",
			InStock:     true,
		},
		{
			Name:        "Headphones",
			Description: "Noise-cancelling wireless headphones",
			Price:       199.99,
			Category:    "Audio",
			InStock:     false,
		},
	}

	for _, product := range products {
		DB.Create(&product)
	}

	log.Println("Database seeded with initial data")
}
