package services

import (
	"errors"
	"time"
	"os"
	"log"

	"github.com/golang-jwt/jwt/v4"
	"github.com/poushiran/internal/database"
	"github.com/poushiran/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var JwtKey = []byte(getEnv("JWT_SECRET", "default_secret_key"))

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Warning: Environment variable %s not set, using default value.", key)
		return fallback
	}
	return value
}

// Claims struct to hold JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Login authenticates a user and returns a JWT token
func Login(username, password string) (string, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("incorrect username or password")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
