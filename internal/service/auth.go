package service

import (
	"github.com/2SSK/jwt/internal/server"
)

type AuthService struct {
	server *server.Server
}

func NewAuthService(s *server.Server) *AuthService {
	return &AuthService{
		server: s,
	}
}
