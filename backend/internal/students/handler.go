package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StudentHandler struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Semester  int    `json:"semester" binding:"required"`
	Branch    string `json:"branch" binding:"required"`
}

func CreateStudentHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input StudentHandler

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		student, err := CreateStudent(pool, input.FirstName, input.LastName, input.Email, input.Phone, input.Semester, input.Branch)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to create student",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, student)
	}
}

func GetStudentHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		students, err := GetAllStudents(pool)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to retrieve students",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, students)
	}
}