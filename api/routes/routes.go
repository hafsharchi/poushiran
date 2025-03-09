package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/poushiran/api/handlers"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *gin.RouterGroup) {
	// Initialize handlers
	productHandler := handlers.NewProductHandler()

	// Product routes
	productRoutes := router.Group("/products")
	{
		productRoutes.GET("", productHandler.GetProducts)
		productRoutes.GET("/:id", productHandler.GetProductByID)
		productRoutes.POST("", productHandler.CreateProduct)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}
}
