package service

import (
	"github.com/2SSK/jwt/internal/lib"
	"github.com/2SSK/jwt/internal/repository"
	"github.com/2SSK/jwt/internal/server"
)

type Services struct {
	Auth       *AuthService
	User       *UserService
	AuthHelper *utils.AuthHelper
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authHelper := utils.NewAuthHelper(repos.User)
	return &Services{
		User:       NewUserService(repos.User, s.Config.Auth.SecretKey),
		Auth:       NewAuthService(s),
		AuthHelper: authHelper,
	}, nil
}
