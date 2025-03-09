package services

import (
	"errors"

	"github.com/poushiran/internal/database"
	"github.com/poushiran/internal/models"
	"gorm.io/gorm"
)

// ProductService handles business logic for products
type ProductService struct{}

// NewProductService creates a new instance of ProductService
func NewProductService() *ProductService {
	return &ProductService{}
}

// GetProducts returns a list of all products
func (s *ProductService) GetProducts() []models.Product {
	var products []models.Product
	database.DB.Find(&products)
	return products
}

// GetProductByID returns a product by ID
func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	result := database.DB.First(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, result.Error
	}
	return &product, nil
}

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(product *models.Product) error {
	result := database.DB.Create(product)
	return result.Error
}

// UpdateProduct updates an existing product
func (s *ProductService) UpdateProduct(id uint, product *models.Product) error {
	// Check if product exists
	existingProduct, err := s.GetProductByID(id)
	if err != nil {
		return err
	}

	// Update product fields
	product.ID = existingProduct.ID
	result := database.DB.Save(product)
	return result.Error
}

// DeleteProduct deletes a product
func (s *ProductService) DeleteProduct(id uint) error {
	// Check if product exists
	_, err := s.GetProductByID(id)
	if err != nil {
		return err
	}

	// Delete product
	result := database.DB.Delete(&models.Product{}, id)
	return result.Error
}
