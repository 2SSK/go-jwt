package handler

import (
	"net/http"

	"github.com/2SSK/jwt/internal/model/user"
	"github.com/2SSK/jwt/internal/service"
	"github.com/2SSK/jwt/internal/validation"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (h *AuthHandler) SignUp(c echo.Context) error {
	var payload user.AddUserPayload
	if err := validation.BindAndValidate(c, &payload); err != nil {
		return err
	}

	response, err := h.userService.SignUp(c.Request().Context(), &payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var payload user.LoginPayload
	if err := validation.BindAndValidate(c, &payload); err != nil {
		return err
	}

	response, err := h.userService.Login(c.Request().Context(), &payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	// Get user ID from context (set by RequireAuth middleware)
	userID, ok := c.Get("user_id").(uuid.UUID)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid authentication")
	}

	// Generate new tokens
	accessToken, refreshToken, err := h.userService.GenerateTokens(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to generate tokens")
	}

	response := user.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.JSON(http.StatusOK, response)
}
