package middleware

import (
	"net/http"

	appErrors "github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/errors"

	"github.com/gin-gonic/gin"
)
func ErrorMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {

		c.Next()

		errs := c.Errors
		if len(errs) == 0 {
			return
		}

		err := errs.Last().Err

		log := GetLogger(c)

		if appErr, ok := err.(*appErrors.AppError); ok {

			log.Warn().
				Err(appErr.Err).
				Str("code", appErr.Code).
				Msg(appErr.Message)

			c.JSON(appErr.Status, gin.H{
				"error": appErr.Message,
				"code" : appErr.Code,
			})
			return
		}

		log.Error().
			Err(err).
			Msg("Unexpected error occurred")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
	}
}