package service

import (
	"context"
	"errors"
	"time"

	"github.com/2SSK/jwt/internal/model/user"
	"github.com/2SSK/jwt/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo  repository.UserRepository
	jwtSecret []byte
}

func NewUserService(userRepo repository.UserRepository, jwtSecret string) *UserService {
	return &UserService{
		userRepo:  userRepo,
		jwtSecret: []byte(jwtSecret),
	}
}

func (s *UserService) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (s *UserService) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *UserService) SignUp(ctx context.Context, payload *user.AddUserPayload) (*user.SignUpResponse, error) {
	// Check if user already exists
	existingUser, err := s.userRepo.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := s.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	// Set default user type to "user" if not provided
	userType := payload.UserType
	if userType == "" {
		userType = "user"
	}

	// Create user
	newUser := &user.User{
		FirstName: &payload.FirstName,
		LastName:  &payload.LastName,
		Password:  &hashedPassword,
		Email:     &payload.Email,
		Phone:     &payload.Phone,
		UserType:  &userType,
	}

	createdUser, err := s.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	// Generate tokens
	accessToken, refreshToken, err := s.generateTokens(createdUser.ID)
	if err != nil {
		return nil, err
	}

	// Update user with tokens (optional, depending on design)
	createdUser.Token = &accessToken
	createdUser.RefreshToken = &refreshToken

	response := &user.SignUpResponse{
		User: user.UserResponse{
			ID:        createdUser.ID,
			FirstName: createdUser.FirstName,
			LastName:  createdUser.LastName,
			Email:     createdUser.Email,
			Phone:     createdUser.Phone,
			UserType:  createdUser.UserType,
			CreatedAt: createdUser.CreatedAt,
			UpdatedAt: createdUser.UpdatedAt,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (s *UserService) Login(ctx context.Context, payload *user.LoginPayload) (*user.LoginResponse, error) {
	// Get user by email
	u, err := s.userRepo.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("invalid credentials")
	}

	// Verify password
	if err := s.VerifyPassword(*u.Password, payload.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate tokens
	accessToken, refreshToken, err := s.generateTokens(u.ID)
	if err != nil {
		return nil, err
	}

	response := &user.LoginResponse{
		User: user.UserResponse{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Phone:     u.Phone,
			UserType:  u.UserType,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (s *UserService) GetUsers(ctx context.Context, limit, offset int) ([]*user.UserResponse, error) {
	users, err := s.userRepo.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []*user.UserResponse
	for _, u := range users {
		responses = append(responses, &user.UserResponse{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Phone:     u.Phone,
			UserType:  u.UserType,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	return responses, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id uuid.UUID, payload *user.UpdateUserPayload) (*user.UserResponse, error) {
	// Get existing user
	existing, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("user not found")
	}

	// Apply updates
	if payload.FirstName != nil {
		existing.FirstName = payload.FirstName
	}
	if payload.LastName != nil {
		existing.LastName = payload.LastName
	}
	if payload.Email != nil {
		// Check email uniqueness if changed
		if *payload.Email != *existing.Email {
			existingByEmail, err := s.userRepo.GetUserByEmail(ctx, *payload.Email)
			if err != nil {
				return nil, err
			}
			if existingByEmail != nil && existingByEmail.ID != id {
				return nil, errors.New("email already in use")
			}
		}
		existing.Email = payload.Email
	}
	if payload.Phone != nil {
		existing.Phone = payload.Phone
	}
	if payload.UserType != nil {
		existing.UserType = payload.UserType
	}

	// Update in DB
	err = s.userRepo.UpdateUser(ctx, existing)
	if err != nil {
		return nil, err
	}

	// Return updated response
	return &user.UserResponse{
		ID:        existing.ID,
		FirstName: existing.FirstName,
		LastName:  existing.LastName,
		Email:     existing.Email,
		Phone:     existing.Phone,
		UserType:  existing.UserType,
		CreatedAt: existing.CreatedAt,
		UpdatedAt: existing.UpdatedAt,
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	// Check if user exists
	existing, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("user not found")
	}

	// Delete
	return s.userRepo.DeleteUser(ctx, id)
}

func (s *UserService) GenerateTokens(ctx context.Context, userID uuid.UUID) (accessToken, refreshToken string, err error) {
	return s.generateTokens(userID)
}

func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*user.UserResponse, error) {
	u, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("user not found")
	}

	return &user.UserResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
		UserType:  u.UserType,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (s *UserService) generateTokens(userID uuid.UUID) (accessToken, refreshToken string, err error) {
	// Access token (short-lived)
	accessClaims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
		"iat":     time.Now().Unix(),
	}
	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = accessTokenObj.SignedString(s.jwtSecret)
	if err != nil {
		return "", "", err
	}

	// Refresh token (long-lived)
	refreshClaims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}
	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = refreshTokenObj.SignedString(s.jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
