package dto

import (
	"github.com/aprilboiz/flight-management/internal/models"
)

type UserResponse struct {
	ID        uint        `json:"id"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Role      models.Role `json:"role"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
