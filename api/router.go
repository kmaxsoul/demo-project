package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kmaxsoul/demo-project/handlers"
)

func SetupRouter(pool *pgxpool.Pool) *gin.Engine {
	var router *gin.Engine = gin.Default()
	router.POST("/todos", handlers.CreateTodoHandler(pool))
	router.GET("/todos", handlers.GetAllTodosHandler(pool))
	router.GET("/todos/:id", handlers.GetTodoByIDHandler(pool))
	router.PUT("/todos/:id", handlers.UpdateTodoHandler(pool))

	return router
}
