package handler

import (
	"net/http"

	"github.com/2SSK/jwt/internal/server"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	server *server.Server
}

func NewHomeHandler(s *server.Server) *HomeHandler {
	return &HomeHandler{server: s}
}

func (h *HomeHandler) ServeHome(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"name":    "Jwt API",
		"version": "v1",
		"health":  "/status",
		"docs":    "/docs",
	})
}
