package students

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func StudentRoutes(router *gin.Engine, pool *pgxpool.Pool) {
	router.POST("/api/students", CreateStudentHandler(pool))
	router.GET("/api/students", GetStudentHandler(pool))
}