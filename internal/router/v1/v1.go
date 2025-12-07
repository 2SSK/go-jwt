package v1

import (
	"github.com/2SSK/jwt/internal/handler"
	"github.com/2SSK/jwt/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterV1Routes(router *echo.Group, handlers *handler.Handlers, middleware *middleware.Middlewares) {
	// Auth routes
	registerAuthRoutes(router, handlers, middleware.Auth)

	// User routes
	registerUserRoutes(router, middleware.Auth, handlers)
}
