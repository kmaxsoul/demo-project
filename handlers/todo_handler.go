package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kmaxsoul/demo-project/repository"
)

type CreateTodoRequest struct {
	Title     string `json:"title" bind:"required"`
	Completed bool   `json:"completed"`
}

type UpdateTodoRequest struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

func CreateTodoHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDInteface, exists := c.Get("userID")
		if !exists {
			c.JSON(500, gin.H{"error": "User ID not found in context"})
			return
		}

		userID := userIDInteface.(string)

		var input CreateTodoRequest

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		todo, err := repository.CreateTodo(pool, input.Title, input.Completed, userID)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, todo)
	}
}

func GetAllTodosHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDInteface, exists := c.Get("userID")
		if !exists {
			c.JSON(500, gin.H{"error": "User ID not found in context"})
			return
		}
		userID := userIDInteface.(string)
		todos, err := repository.GetAllTodos(pool, userID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, todos)
	}

}

func GetTodoByIDHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDInteface, exists := c.Get("userID")
		if !exists {
			c.JSON(500, gin.H{"error": "User ID not found in context"})
			return
		}

		userID := userIDInteface.(string)
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		todo, err := repository.GetTodoByID(pool, id, userID)
		if err != nil {

			if err == pgx.ErrNoRows {
				c.JSON(404, gin.H{"error": "Todo not found"})
				return
			}

			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, todo)
	}
}

func UpdateTodoHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDInteface, exists := c.Get("userID")
		if !exists {
			c.JSON(500, gin.H{"error": "User ID not found in context"})
			return
		}

		userID := userIDInteface.(string)

		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		var input UpdateTodoRequest

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if input.Title == nil && input.Completed == nil {
			c.JSON(400, gin.H{"error": "No fields to update"})
			return
		}

		existing, err := repository.GetTodoByID(pool, id, userID)
		if err != nil {
			if err == pgx.ErrNoRows {
				c.JSON(404, gin.H{"error": "Todo not found"})
				return
			}
		}

		title := existing.Title
		if input.Title != nil {
			title = *input.Title
		}

		var completed bool
		completed = existing.Completed
		if input.Completed != nil {
			completed = *input.Completed
		}

		todo, err := repository.UpdateTodo(pool, id, title, completed, userID)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, todo)
	}
}

func DeleteTodoHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDInteface, exists := c.Get("userID")
		if !exists {
			c.JSON(500, gin.H{"error": "User ID not found in context"})
			return
		}

		userID := userIDInteface.(string)

		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		err = repository.DeleteTodo(pool, id, userID)
		if err != nil {
			if err == pgx.ErrNoRows {
				c.JSON(404, gin.H{"error": "Todo not found"})
				return
			}

			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Todo deleted successfully"})
	}
}
