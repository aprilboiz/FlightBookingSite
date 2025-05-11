package handlers

import (
	"net/http"

	"github.com/aprilboiz/flight-management/internal/dto"
	e "github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type userHandler struct {
	userService service.UserService
	logger      *zap.Logger
}

func NewUserHandler(userService service.UserService, logger *zap.Logger) UserHandler {
	return &userHandler{
		userService: userService,
		logger:      logger,
	}
}

// Register handles user registration
//	@Summary		Register a new user
//	@Description	Register a new user with username, password, and email
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.RegisterRequest	true	"User registration information"
//	@Success		200		{object}	dto.AuthResponse
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		409		{object}	dto.ErrorResponse
//	@Router			/auth/register [post]
func (h *userHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(e.BadRequest("Invalid request body", err))
		return
	}

	response, err := h.userService.Register(req)
	if err != nil {
		switch err {
		case service.ErrUserExists:
			_ = c.Error(e.NewAppError(e.CONFLICT, "User already exists", nil))
		default:
			h.logger.Error("Failed to register user", zap.Error(err))
			_ = c.Error(e.Internal("Failed to register user", err))
		}
		return
	}

	c.JSON(http.StatusOK, response)
}

// Login handles user login
//	@Summary		Login user
//	@Description	Login with username and password
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.LoginRequest	true	"User login credentials"
//	@Success		200		{object}	dto.AuthResponse
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Router			/auth/login [post]
func (h *userHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(e.BadRequest("Invalid request body", err))
		return
	}

	response, err := h.userService.Login(req)
	if err != nil {
		switch err {
		case service.ErrInvalidCredentials:
			_ = c.Error(e.NewAppError(e.UNAUTHORIZED, "Invalid credentials", nil))
		default:
			h.logger.Error("Failed to login user", zap.Error(err))
			_ = c.Error(e.Internal("Failed to login", err))
		}
		return
	}

	c.JSON(http.StatusOK, response)
}
