package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kmaxsoul/demo-project/config"
	"github.com/kmaxsoul/demo-project/handlers"
	"github.com/kmaxsoul/demo-project/middleware"
)

func SetupRouter(pool *pgxpool.Pool, cfg *config.Config) *gin.Engine {
	var router *gin.Engine = gin.Default()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", handlers.CreateUserHandler(pool))
		authGroup.POST("/login", handlers.LoginUserHandler(pool, cfg))
	}

	todoGroup := router.Group("/todos")
	todoGroup.Use(middleware.AuthMiddleware(cfg))
	{
		todoGroup.POST("", handlers.CreateTodoHandler(pool))
		todoGroup.GET("", handlers.GetAllTodosHandler(pool))
		todoGroup.GET("/:id", handlers.GetTodoByIDHandler(pool))
		todoGroup.PUT("/:id", handlers.UpdateTodoHandler(pool))
		todoGroup.DELETE("/:id", handlers.DeleteTodoHandler(pool))
	}

	return router

}
