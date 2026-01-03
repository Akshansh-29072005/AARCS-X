package main

import (
	"log"

	"github.com/Akshansh-29072005/AARCS-X/backend/internals/config"
	"github.com/Akshansh-29072005/AARCS-X/backend/internals/database"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	var cfg *config.Config
	var err error
	cfg, err = config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer pool.Close()

	log.Println("Starting AARCS-X API server...")

	var router *gin.Engine = gin.Default()

	router.SetTrustedProxies(nil)

	router.GET("/", func(c *gin.Context){
		//map[string]interface{}
		c.JSON(200, gin.H{
			"message" : "AARCS-X API is running!",
			"status" : "success",
			"database_connected" : true,
		})
	})

	router.Run(":" + cfg.Port)
}