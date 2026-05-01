package institutes

import (
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *Handler, rp middleware.RoleProvider) {

	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())
	v1.Use(middleware.RequireRole(rp, "user"))

	// Institution Creating Route
	v1.POST("/institutions", h.CreateInstitute)
	// Institution Info Getting Route
	v1.GET("/institutions", h.Read)


	superUser := r.Group("/api/v1")
	superUser.Use(middleware.AuthMiddleware())
	superUser.Use(middleware.RequireRole(rp, "super_user"))

	{
		// Super User Routes for managing institutions
	}

	ownerOnly := r.Group("/api/v1")
	ownerOnly.Use(middleware.AuthMiddleware())
	ownerOnly.Use(middleware.RequireRole(rp, "institution_owner"))

	{
		// Make Admin Route, only institution owners can promote admins
		ownerOnly.POST("/institutions/make-admin", h.MakeAdmin)
	}

	staff := r.Group("/api/v1")
	staff.Use(middleware.AuthMiddleware())
	staff.Use(middleware.RequireRole(rp, "institution_owner", "institution_admin"))
	{
		// Institution Info Getting by ID Route
		staff.GET("/institutions/:id", h.ReadByID)
	}
}
