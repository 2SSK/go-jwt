package utils

import (
	"context"
	"errors"

	"github.com/2SSK/jwt/internal/repository"
	"github.com/google/uuid"
)

type AuthHelper struct {
	userRepo repository.UserRepository
}

func NewAuthHelper(userRepo repository.UserRepository) *AuthHelper {
	return &AuthHelper{userRepo: userRepo}
}

// CheckUserType verifies if a user has the required user type
func (h *AuthHelper) CheckUserType(ctx context.Context, userID uuid.UUID, requiredType string) error {
	user, err := h.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	if user.UserType == nil || *user.UserType != requiredType {
		return errors.New("insufficient permissions")
	}

	return nil
}

// MatchUserTypeToUserId checks if the user has the specified type (admin or user)
// This is just checking if everything is fine
func (h *AuthHelper) MatchUserTypeToUserId(ctx context.Context, userID uuid.UUID, userType string) (bool, error) {
	err := h.CheckUserType(ctx, userID, userType)
	if err != nil {
		return false, err
	}
	return true, nil
}
