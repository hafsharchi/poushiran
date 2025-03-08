package services

import (
	"github.com/poushiran/internal/models"
)

// ProductService handles business logic for products
type ProductService struct{}

// NewProductService creates a new instance of ProductService
func NewProductService() *ProductService {
	return &ProductService{}
}

// GetProducts returns a list of products
func (s *ProductService) GetProducts() []models.Product {
	// Mock data for demonstration
	return []models.Product{
		{
			ID:          "1",
			Name:        "Laptop",
			Description: "High-performance laptop with 16GB RAM",
			Price:       999.99,
			Category:    "Electronics",
			InStock:     true,
		},
		{
			ID:          "2",
			Name:        "Smartphone",
			Description: "Latest smartphone with 128GB storage",
			Price:       699.99,
			Category:    "Electronics",
			InStock:     true,
		},
		{
			ID:          "3",
			Name:        "Headphones",
			Description: "Noise-cancelling wireless headphones",
			Price:       199.99,
			Category:    "Audio",
			InStock:     false,
		},
	}
} 