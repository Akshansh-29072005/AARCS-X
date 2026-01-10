package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(router *gin.Engine, pool *pgxpool.Pool) {

	//Health Check Route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": 		  "AARCS-X API is running!",
			"status":   	  "success",
			"database_connected": true,
		})
	})
}