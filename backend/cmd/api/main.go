package main

import (
	"log"

	"github.com/Akshansh-29072005/AARCS-X/backend/internals/config"
	"github.com/Akshansh-29072005/AARCS-X/backend/internals/database"
	"github.com/Akshansh-29072005/AARCS-X/backend/internals/handlers"
	"github.com/gin-contrib/cors"
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

	// Configuring CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))

	router.GET("/", func(c *gin.Context) {
		//map[string]interface{}
		c.JSON(200, gin.H{
			"message":            "AARCS-X API is running!",
			"status":             "success",
			"database_connected": true,
		})
	})

	router.POST("/api/students", handlers.CreateStudentHandler(pool))

	router.Run(":" + cfg.Port)
}
