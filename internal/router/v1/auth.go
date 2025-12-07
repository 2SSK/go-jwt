package v1

import (
	"github.com/2SSK/jwt/internal/handler"
	"github.com/2SSK/jwt/internal/middleware"
	"github.com/labstack/echo/v4"
)

func registerAuthRoutes(r *echo.Group, handlers *handler.Handlers, authMiddleware *middleware.AuthMiddleware) {
	// Auth routes
	auth := r.Group("/auth")

	// Auth Operations
	auth.POST("/signup", handlers.Auth.SignUp)                                      // User Signup
	auth.POST("/login", handlers.Auth.Login)                                        // User Login
	auth.POST("/refresh", handlers.Auth.RefreshToken, authMiddleware.RequireAuth()) // Refresh Token
}
