package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kmaxsoul/demo-project/api"
	"github.com/kmaxsoul/demo-project/config"
	"github.com/kmaxsoul/demo-project/database"
)

func main() {
	var cfg *config.Config
	var err error
	cfg, err = config.Load()
	if err != nil {
		log.Printf("Failed to load configuration")
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Printf("Failed to connect to database")
	}

	defer pool.Close()

	var router *gin.Engine = api.SetupRouter(pool, cfg)

	router.Run(":" + cfg.Port)
}
