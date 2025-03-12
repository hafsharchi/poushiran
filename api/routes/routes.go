package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/poushiran/api/handlers"
	"github.com/poushiran/internal/services"
)

// JWTAuthMiddleware checks for a valid JWT token
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		claims := &services.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return services.JwtKey, nil
		})

		if err != nil {
			c.JSON(401, gin.H{"error": "Error parsing token: " + err.Error()})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(401, gin.H{"error": "Token is invalid"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *gin.RouterGroup) {
	// Initialize handlers
	productHandler := handlers.NewProductHandler()

	// Product routes
	productRoutes := router.Group("/products")
	{
		productRoutes.GET("", productHandler.GetProducts)
		productRoutes.GET("/:id", productHandler.GetProductByID)
		productRoutes.POST("", JWTAuthMiddleware(), productHandler.CreateProduct)
		productRoutes.PUT("/:id", JWTAuthMiddleware(), productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", JWTAuthMiddleware(), productHandler.DeleteProduct)
	}

	// User routes
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/login", handlers.LoginHandler)
	}
}
