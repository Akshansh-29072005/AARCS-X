package teachers

import(
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TeacherRoutes(router *gin.Engine, pool *pgxpool.Pool) {

	// Create Teacher Route
	router.POST("/api/teachers", CreateTeacherHandler(pool))

	// Get Teachers Route
	// router.GET("/api/teachers", GetTeacherHandler(pool))
}