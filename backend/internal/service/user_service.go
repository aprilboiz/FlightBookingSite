package service

import (
	"errors"
	"time"

	"github.com/aprilboiz/flight-management/internal/dto"
	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/aprilboiz/flight-management/internal/repository"
	"github.com/aprilboiz/flight-management/pkg/auth"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(req dto.RegisterRequest) (*dto.AuthResponse, error) {
	// Check if user already exists
	if _, err := s.userRepo.GetByUsername(req.Username); err == nil {
		return nil, ErrUserExists
	}
	if _, err := s.userRepo.GetByEmail(req.Email); err == nil {
		return nil, ErrUserExists
	}

	user := &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Role:     models.RoleStaff, // Default role
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		},
	}, nil
}

func (s *userService) Login(req dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := user.CheckPassword(req.Password); err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := auth.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		},
	}, nil
}
