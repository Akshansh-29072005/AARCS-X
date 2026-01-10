package teachers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TeacherHandler struct {
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Phone      string `json:"phone" binding:"required"`
	Department string `json:"department" binding:"required"`
	Designation string `json:"designation" binding:"required"`
}

func CreateTeacherHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input TeacherHandler

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		teacher, err := CreateTeacher(pool, input.FirstName, input.LastName, input.Email, input.Phone, input.Department, input.Designation)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to create teacher",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, teacher)
	}
}
