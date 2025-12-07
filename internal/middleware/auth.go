package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/2SSK/jwt/internal/server"
	"github.com/2SSK/jwt/internal/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	server   *server.Server
	services *service.Services
}

func NewAuthMiddleware(s *server.Server, services *service.Services) *AuthMiddleware {
	return &AuthMiddleware{server: s, services: services}
}

func (auth *AuthMiddleware) RequireRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// First, require authentication
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization header format")
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return []byte(auth.server.Config.Auth.SecretKey), nil
			})

			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token claims")
			}

			userIDStr, ok := claims["user_id"].(string)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid user ID in token")
			}

			userID, err := uuid.Parse(userIDStr)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid user ID format")
			}

			// Check user role using AuthHelper
			err = auth.services.AuthHelper.CheckUserType(c.Request().Context(), userID, role)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions")
			}

			// Set user ID in context for handlers to use
			c.Set("user_id", userID)

			return next(c)
		}
	}
}

func (auth *AuthMiddleware) RequireAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization header format")
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return []byte(auth.server.Config.Auth.SecretKey), nil
			})

			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token claims")
			}

			userIDStr, ok := claims["user_id"].(string)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid user ID in token")
			}

			userID, err := uuid.Parse(userIDStr)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid user ID format")
			}

			// Set user ID in context for handlers to use
			c.Set("user_id", userID)

			return next(c)
		}
	}
}
