package main

import (
	"log/slog"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/auth"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/config"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/departments"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/institutes"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/database"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/logger"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/server"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/utlis"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/semesters"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/students"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/subjects"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/teachers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	zerolog_logger := logger.NewLogger()
	
	var cfg *config.Config
	var err error
	cfg, err = config.Load()
	if err != nil {
		zerolog_logger.Error().Err(err).Msg("Failed to load configuration")
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)

	if err != nil {
		zerolog_logger.Fatal().Err(err).Msg("Failed to connect to database")
	}

	defer pool.Close()

	zerolog_logger.Info().Msg("Starting AARCS-X API server...")

	var router = gin.Default()

	gin.SetMode(cfg.GinMode)

	err = router.SetTrustedProxies(nil)
	if err != nil {
		zerolog_logger.Fatal().Err(err).Msg("Failed to set trusted proxies")
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

	// Setting JWT Secret
	utlis.SetJWTSecret(cfg.JWTSecret)

	var (
		// Institution Repository
		institutionRepository = institutes.NewRepository(pool)

		// Institution Service
		institutionService = institutes.NewService(institutionRepository)

		// Institution Handler
		institutionHandler = institutes.NewHandler(institutionService)

		// Department Repository
		departmentRepository = departments.NewRepository(pool)

		// Department Service
		departmentService = departments.NewService(departmentRepository)

		// Department Handler
		departmentHandler = departments.NewHandler(departmentService)

		// Semester Repository
		semestersRepository = semesters.NewRepository(pool)

		// Semester Service
		semestersService = semesters.NewService(semestersRepository)

		// Semester Handler
		semestersHandler = semesters.NewHandler(semestersService)

		// Subject Repository
		subjectRepository = subjects.NewRepository(pool)

		// Subject Service
		subjectsService = subjects.NewService(subjectRepository)

		// Subject handler
		subjectHandler = subjects.NewHandler(subjectsService)

		// Teacher Repository
		teacherRepository = teachers.NewRepository(pool)

		// Teacher Service
		teacherService = teachers.NewService(teacherRepository)

		// Teacher Handler
		teacherHandler = teachers.NewHandler(teacherService)

		// Student Repository
		studentRepository = students.NewRepository(pool)

		// Student Service
		studentService = students.NewService(studentRepository)

		// Student Handler
		studentHandler = students.NewHandler(studentService)

		// Auth Repository
		authRepository = auth.NewRepository(pool)

		// Auth Service
		authService = auth.NewService(authRepository)

		// Auth Handler
		authHandler = auth.NewHandler(authService, institutionService)
	)

	// Auth Routes
	auth.RegisteredRoutes(router, authHandler)

	// Institution Routes
	institutes.RegisterRoutes(router, institutionHandler)

	// Department Routes
	departments.RegisterRoutes(router, departmentHandler)

	// Semester Routes
	semesters.RegisterRoutes(router, semestersHandler)

	// Subject Routes
	subjects.RegisterRoutes(router, subjectHandler)

	// Teacher Routes
	teachers.RegisteredRoutes(router, teacherHandler)

	// Student Routes
	students.RegisterRoutes(router, studentHandler)

	if err = router.Run(":" + cfg.Port); err != nil {
		zerolog_logger.Fatal().Err(err).Msg("Failed to start server")
	}
}
