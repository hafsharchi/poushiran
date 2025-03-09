package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/poushiran/api/routes"
	"github.com/poushiran/config"
	_ "github.com/poushiran/docs"
	"github.com/poushiran/internal/database"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Poushiran API
// @version         1.0
// @description     A sample API for Poushiran project
// @host            localhost:8080
// @BasePath        /api/v1

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	err := database.Initialize(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Setup API routes
	apiV1 := r.Group("/api/v1")
	routes.SetupRoutes(apiV1)

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
