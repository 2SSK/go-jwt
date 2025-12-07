package handler

import (
	"github.com/2SSK/jwt/internal/server"
	"github.com/2SSK/jwt/internal/service"
)

type Handlers struct {
	Health  *HealthHandler
	OpenAPI *OpenAPIHandler
	Home    *HomeHandler
	Auth    *AuthHandler
	User    *UserHandler
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{
		Health:  NewHealthHandler(s),
		OpenAPI: NewOpenAPIHandler(s),
		Home:    NewHomeHandler(s),
		Auth:    NewAuthHandler(services.User),
		User:    NewUserHandler(services.User, services.AuthHelper),
	}
}
