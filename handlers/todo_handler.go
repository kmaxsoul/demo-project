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

func CreateTodoHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CreateTodoRequest

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		todo, err := repository.CreateTodo(pool, input.Title, input.Completed)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, todo)
	}
}

func GetAllTodosHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		todos, err := repository.GetAllTodos(pool)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, todos)
	}

}

func GetTodoByIDHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		todo, err := repository.GetTodoByID(pool, id)
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
