package middleware

import (
	"github.com/2SSK/jwt/internal/server"
	"github.com/2SSK/jwt/internal/service"
)

type Middlewares struct {
	Global          *GlobalMiddlewares
	RateLimit       *RateLimitMiddleware
	ContextEnhancer *ContextEnhancer
	Auth            *AuthMiddleware
}

func NewMiddlewares(s *server.Server, services *service.Services) *Middlewares {

	return &Middlewares{
		Global:          NewGlobalMiddlewares(s),
		RateLimit:       NewRateLimitMiddleware(s),
		ContextEnhancer: NewContextEnhancer(s),
		Auth:            NewAuthMiddleware(s, services),
	}
}
