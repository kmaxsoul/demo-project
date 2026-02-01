package handlers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kmaxsoul/demo-project/models"
	"github.com/kmaxsoul/demo-project/repository"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUserHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {

		var registerRequest RegisterRequest

		if err := c.BindJSON(&registerRequest); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		if len(registerRequest.Password) < 8 {
			c.JSON(400, gin.H{"error": "Password must be at least 8 characters long"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to hash password"})
			return
		}

		user := models.User{
			Email:    registerRequest.Email,
			Password: string(hashedPassword),
		}

		createuser, err := repository.CreateUser(pool, &user)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique constraint") {
				c.JSON(400, gin.H{"error": "Email already in use"})
				return
			}
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"message": "User created successfully", "user": createuser})
	}
}
