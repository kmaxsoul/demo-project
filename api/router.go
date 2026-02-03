package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kmaxsoul/demo-project/config"
	"github.com/kmaxsoul/demo-project/handlers"
)

func SetupRouter(pool *pgxpool.Pool, cfg *config.Config) *gin.Engine {
	var router *gin.Engine = gin.Default()

	todoGroup := router.Group("/todos")
	{
		todoGroup.POST("", handlers.CreateTodoHandler(pool))
		todoGroup.GET("", handlers.GetAllTodosHandler(pool))
		todoGroup.GET("/:id", handlers.GetTodoByIDHandler(pool))
		todoGroup.PUT("/:id", handlers.UpdateTodoHandler(pool))
		todoGroup.DELETE("/:id", handlers.DeleteTodoHandler(pool))
	}

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", handlers.CreateUserHandler(pool))
		authGroup.POST("/login", handlers.LoginUserHandler(pool, cfg))
	}

	return router

}
