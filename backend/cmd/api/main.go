package main

import (
	"log"
	"log/slog"

	// "github.com/Akshansh-29072005/AARCS-X/backend/internal/auth"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/config"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/database"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/server"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/students"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/teachers"
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

	var router = gin.Default()

	err = router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal("Failed to set trusted proxies:", err)
	}

	// Configuring CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))

	var logger = slog.Default()

	//Auth Routes
	//auth.RegisterRoutes(router, authHandler)

	// Student Service Enabling

	//Server Status Routes
	server.RegisterRoutes(router, pool, logger)

	// Student Repository
	var studentRepository = students.NewRepository(pool)

	// Student Service
	var studentService = students.NewService(studentRepository)

	// Student Handler
	var studentHandler = students.NewHandler(studentService)

	// Student Routes
	students.RegisterRoutes(router, studentHandler)

	// Teacher Service Enabling

	// Teacher Routes
	teachers.TeacherRoutes(router, pool)

	if err = router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
