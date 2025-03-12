package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poushiran/internal/services"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary User Login
// @Description Authenticate user and return a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param loginRequest body LoginRequest true "Login request payload"
// @Router /users/login [post]
func LoginHandler(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Authenticate user and generate JWT token
	token, err := services.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
