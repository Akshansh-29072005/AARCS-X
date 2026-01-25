package main

import (
	"log"
	"log/slog"

	// "github.com/Akshansh-29072005/AARCS-X/backend/internal/auth"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/config"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/departments"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/institutes"
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

	//Server Status Routes
	server.RegisterRoutes(router, pool, logger)

	// Enabling Auth Routes

	// Auth Repository
	// var authRepository = auth.NewRepository(pool)

	/*
		Institution Service Enabling
	*/

	// Institution Repository
	var institutionRepository = institutes.NewRepository(pool)

	// Institution Service
	var institutionService = institutes.NewService(institutionRepository)

	// Institution Handler
	var institutionHandler = institutes.NewHandler(institutionService)

	// Institution Routes
	institutes.RegisterRoutes(router, institutionHandler)

	/*
		Department Service Enabling
	*/

	// Department Repository
	var departmentRepository = departments.NewRepository(pool)

	// Department Service
	var departmentService = departments.NewService(departmentRepository)

	// Department Handler
	var departmentHandler = departments.NewHandler(departmentService)

	// Department Routes
	departments.RegisterRoutes(router, departmentHandler)

	/*
		Student Service Enabling
	*/

	// Student Repository
	var studentRepository = students.NewRepository(pool)

	// Student Service
	var studentService = students.NewService(studentRepository)

	// Student Handler
	var studentHandler = students.NewHandler(studentService)

	// Student Routes
	students.RegisterRoutes(router, studentHandler)

	/*
		Teacher Service Enabling
	*/

	// Teacher Repository
	var teacherRepository = teachers.NewRepository(pool)

	// Teacher Service
	var teacherService = teachers.NewService(teacherRepository)

	// Teacher Handler
	var teacherHandler = teachers.NewHandler(teacherService)

	// Teacher Routes
	teachers.RegisteredRoutes(router, teacherHandler)

	if err = router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
