package v1

import (
	"github.com/2SSK/jwt/internal/handler"
	"github.com/2SSK/jwt/internal/middleware"
	"github.com/labstack/echo/v4"
)

func registerUserRoutes(r *echo.Group, auth *middleware.AuthMiddleware, handlers *handler.Handlers) {
	// Users routes
	users := r.Group("/users")
	users.Use(auth.RequireRole("admin")) // Admin only

	// Users Operations
	users.GET("", handlers.User.GetUsers) // List Users

	// Admin Operations
	admin := r.Group("/user")
	admin.Use(auth.RequireRole("admin"))                // Admin only
	admin.GET("/:user_id", handlers.User.GetUserByID)   // Get User by ID
	admin.PUT("/:user_id", handlers.User.UpdateUser)    // Update User
	admin.DELETE("/:user_id", handlers.User.DeleteUser) // Delete User
}
