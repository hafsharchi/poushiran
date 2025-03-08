package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poushiran/internal/services"
)

// ProductHandler handles HTTP requests related to products
type ProductHandler struct {
	productService *services.ProductService
}

// NewProductHandler creates a new instance of ProductHandler
func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		productService: services.NewProductService(),
	}
}

// GetProducts godoc
// @Summary      Get all products
// @Description  Returns a list of all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Product
// @Router       /products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	products := h.productService.GetProducts()
	c.JSON(http.StatusOK, products)
} 