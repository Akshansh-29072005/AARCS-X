package server

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

const (
	// Health check sampling interval
	healthCheckInterval = 10 * time.Second
	// Memory conversion from bytes to MB
	bytesPerMB = 1024 * 1024
	// CPU sampling duration
	cpuSampleDuration = 1 * time.Second
)

// HealthMetrics holds current system health metrics
type HealthMetrics struct {
	mu            sync.RWMutex
	CPUUsage      float64
	TotalMemory   uint64
	UsedMemory    uint64
	LastUpdated   time.Time
	DatabaseReady bool
}

// HealthService manages system health monitoring
type HealthService struct {
	metrics *HealthMetrics
	logger  *slog.Logger
	pool    *pgxpool.Pool
	done    chan struct{}
}

// NewHealthService creates a new health service instance
func NewHealthService(pool *pgxpool.Pool, logger *slog.Logger) *HealthService {
	return &HealthService{
		metrics: &HealthMetrics{},
		logger:  logger,
		pool:    pool,
		done:    make(chan struct{}),
	}
}

// Start begins the background health monitoring goroutine
func (hs *HealthService) Start(ctx context.Context) {
	go hs.monitorHealth(ctx)
	hs.logger.Info("health service started")
}

// Stop gracefully stops the health monitoring
func (hs *HealthService) Stop() {
	close(hs.done)
	hs.logger.Info("health service stopped")
}

// monitorHealth continuously updates system metrics
func (hs *HealthService) monitorHealth(ctx context.Context) {
	ticker := time.NewTicker(healthCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-hs.done:
			return
		case <-ticker.C:
			hs.updateMetrics()
		}
	}
}

// updateMetrics refreshes current system metrics
func (hs *HealthService) updateMetrics() {
	// Check database connectivity
	dbReady := hs.checkDatabase()

	// Get memory stats
	memStats, err := memory.Get()
	if err != nil {
		hs.logger.Warn("failed to get memory stats", "error", err)
		return
	}

	// Get CPU usage
	cpuUsage, err := hs.calculateCPUUsage()
	if err != nil {
		hs.logger.Warn("failed to calculate cpu usage", "error", err)
		return
	}

	// Update metrics atomically
	hs.metrics.mu.Lock()
	hs.metrics.CPUUsage = cpuUsage
	hs.metrics.TotalMemory = memStats.Total / bytesPerMB
	hs.metrics.UsedMemory = memStats.Used / bytesPerMB
	hs.metrics.DatabaseReady = dbReady
	hs.metrics.LastUpdated = time.Now()
	hs.metrics.mu.Unlock()

	hs.logger.Debug("metrics updated",
		"cpu_usage", cpuUsage,
		"memory_used_mb", hs.metrics.UsedMemory,
		"memory_total_mb", hs.metrics.TotalMemory,
	)
}

// calculateCPUUsage returns CPU usage percentage
func (hs *HealthService) calculateCPUUsage() (float64, error) {
	before, err := cpu.Get()
	if err != nil {
		return 0, err
	}

	time.Sleep(cpuSampleDuration)

	after, err := cpu.Get()
	if err != nil {
		return 0, err
	}

	totalDelta := after.Total - before.Total
	idleDelta := after.Idle - before.Idle

	if totalDelta == 0 {
		return 0, nil
	}

	usage := float64(totalDelta-idleDelta) / float64(totalDelta) * 100.0
	return usage, nil
}

// checkDatabase verifies database connectivity
func (hs *HealthService) checkDatabase() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := hs.pool.Ping(ctx)
	return err == nil
}

// GetMetrics returns current metrics in a thread-safe manner
func (hs *HealthService) GetMetrics() HealthMetrics {
	hs.metrics.mu.RLock()
	defer hs.metrics.mu.RUnlock()
	return *hs.metrics
}

// HealthResponse structures for API endpoints
type HealthResponse struct {
	Status            string    `json:"status"`
	Message           string    `json:"message"`
	DatabaseConnected bool      `json:"database_connected"`
	Timestamp         time.Time `json:"timestamp"`
}

type SystemMetricsResponse struct {
	Status        string    `json:"status"`
	CPUUsage      float64   `json:"cpu_usage_percent"`
	TotalMemory   uint64    `json:"total_memory_mb"`
	MemoryUsage   uint64    `json:"memory_usage_mb"`
	MemoryPercent float64   `json:"memory_usage_percent"`
	DatabaseReady bool      `json:"database_ready"`
	LastUpdated   time.Time `json:"last_updated"`
}

// RegisterRoutes sets up all API routes
func RegisterRoutes(router *gin.Engine, pool *pgxpool.Pool, logger *slog.Logger) {
	// Initialize health service
	healthService := NewHealthService(pool, logger)
	healthService.Start(context.Background())

	// Health check endpoint - simple connectivity check
	router.GET("/api/v1/system/health", func(c *gin.Context) {
		response := HealthResponse{
			Status:            "ok",
			Message:           "AARCS-X API is running",
			DatabaseConnected: healthService.metrics.DatabaseReady,
			Timestamp:         time.Now(),
		}
		c.JSON(200, response)
	})

	// System metrics endpoint - detailed system statistics
	router.GET("/api/v1/system/metrics", func(c *gin.Context) {
		metrics := healthService.GetMetrics()

		memoryPercent := float64(0)
		if metrics.TotalMemory > 0 {
			memoryPercent = (float64(metrics.UsedMemory) / float64(metrics.TotalMemory)) * 100.0
		}

		response := SystemMetricsResponse{
			Status:        "ok",
			CPUUsage:      metrics.CPUUsage,
			TotalMemory:   metrics.TotalMemory,
			MemoryUsage:   metrics.UsedMemory,
			MemoryPercent: memoryPercent,
			DatabaseReady: metrics.DatabaseReady,
			LastUpdated:   metrics.LastUpdated,
		}
		c.JSON(200, response)
	})

	logger.Info("routes registered successfully")
}
