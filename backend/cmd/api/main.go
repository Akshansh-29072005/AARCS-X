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
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/semesters"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/students"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/subjects"
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

	//gin.SetMode(gin.ReleaseMode)

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

	var (
		// Institution Repository
		institutionRepository = institutes.NewRepository(pool)

		// Institution Service
		institutionService = institutes.NewService(institutionRepository)

		// Institution Handler
		institutionHandler = institutes.NewHandler(institutionService)
	)

	// Institution Routes
	institutes.RegisterRoutes(router, institutionHandler)

	/*
		Department Service Enabling
	*/

	var (
		// Department Repository
		departmentRepository = departments.NewRepository(pool)

		// Department Service
		departmentService = departments.NewService(departmentRepository)

		// Department Handler
		departmentHandler = departments.NewHandler(departmentService)
	)

	// Department Routes
	departments.RegisterRoutes(router, departmentHandler)

	/*
		Semester Service Enabling
	*/

	var (
		// Semester Repository
		semestersRepository = semesters.NewRepository(pool)

		// Semester Service
		semestersService = semesters.NewService(semestersRepository)

		// Semester Handler
		semestersHandler = semesters.NewHandler(semestersService)
	)

	// Semester Routes
	semesters.RegisterRoutes(router, semestersHandler)

	/*
		Subject Service Enabling
	*/

	var (
		// Subject Repository
		subjectRepository = subjects.NewRepository(pool)

		// Subject Service
		subjectsService = subjects.NewService(subjectRepository)

		// Subject handler
		subjectHandler = subjects.NewHandler(subjectsService)
	)

	// Subject Routes
	subjects.RegisterRoutes(router, subjectHandler)

	/*
		Teacher Service Enabling
	*/

	var (
		// Teacher Repository
		teacherRepository = teachers.NewRepository(pool)

		// Teacher Service
		teacherService = teachers.NewService(teacherRepository)

		// Teacher Handler
		teacherHandler = teachers.NewHandler(teacherService)
	)

	// Teacher Routes
	teachers.RegisteredRoutes(router, teacherHandler)

	/*
		Student Service Enabling
	*/

	var (
		// Student Repository
		studentRepository = students.NewRepository(pool)

		// Student Service
		studentService = students.NewService(studentRepository)

		// Student Handler
		studentHandler = students.NewHandler(studentService)
	)

	// Student Routes
	students.RegisterRoutes(router, studentHandler)

	if err = router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
