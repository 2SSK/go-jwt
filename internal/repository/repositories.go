package repository

import (
	"context"

	"github.com/2SSK/jwt/internal/model/user"
	"github.com/2SSK/jwt/internal/server"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *user.User) (*user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*user.User, error)
	GetUsers(ctx context.Context, limit, offset int) ([]*user.User, error)
	UpdateUser(ctx context.Context, user *user.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type Repositories struct {
	User UserRepository
}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{
		User: NewUserRepository(s.DB.Pool),
	}
}
