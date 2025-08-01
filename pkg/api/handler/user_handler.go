package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"go_ex01/pkg/api/service"
	"go_ex01/pkg/api/vo"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) RegisterRoutes(g *echo.Group) {
	users := g.Group("/users")
	users.GET("", h.GetUsers)
	users.GET("/:id", h.GetUser)
	users.POST("", h.CreateUser)
	users.PUT("/:id", h.UpdateUser)
	users.DELETE("/:id", h.DeleteUser)
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	slog.Debug("Getting all users")

	users := h.userService.GetAllUsers()
	response := vo.UsersResponse{
		Users: users,
		Total: len(users),
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Warn("Invalid user ID", "id", c.Param("id"), "error", err)
		return c.JSON(http.StatusBadRequest, vo.ErrorResponse{
			Error:   "invalid_id",
			Message: "User ID must be a valid integer",
		})
	}

	slog.Debug("Getting user", "id", id)

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, vo.ErrorResponse{
			Error:   "user_not_found",
			Message: "User not found",
		})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req vo.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		slog.Warn("Invalid request body", "error", err)
		return c.JSON(http.StatusBadRequest, vo.ErrorResponse{
			Error:   "invalid_request",
			Message: "Invalid request body",
		})
	}

	if req.Name == "" || req.Email == "" {
		return c.JSON(http.StatusBadRequest, vo.ErrorResponse{
			Error:   "validation_error",
			Message: "Name and email are required",
		})
	}

	slog.Debug("Creating user", "name", req.Name, "email", req.Email)

	user, err := h.userService.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusConflict, vo.ErrorResponse{
			Error:   "creation_failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Warn("Invalid user ID", "id", c.Param("id"), "error", err)
		return c.JSON(http.StatusBadRequest, vo.ErrorResponse{
			Error:   "invalid_id",
			Message: "User ID must be a valid integer",
		})
	}

	var req vo.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		slog.Warn("Invalid request body", "error", err)
		return c.JSON(http.StatusBadRequest, vo.ErrorResponse{
			Error:   "invalid_request",
			Message: "Invalid request body",
		})
	}

	slog.Debug("Updating user", "id", id, "name", req.Name, "email", req.Email)

	user, err := h.userService.UpdateUser(id, req)
	if err != nil {
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNotFound, vo.ErrorResponse{
				Error:   "user_not_found",
				Message: "User not found",
			})
		}
		return c.JSON(http.StatusConflict, vo.ErrorResponse{
			Error:   "update_failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Warn("Invalid user ID", "id", c.Param("id"), "error", err)
		return c.JSON(http.StatusBadRequest, vo.ErrorResponse{
			Error:   "invalid_id",
			Message: "User ID must be a valid integer",
		})
	}

	slog.Debug("Deleting user", "id", id)

	err = h.userService.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, vo.ErrorResponse{
			Error:   "user_not_found",
			Message: "User not found",
		})
	}

	return c.NoContent(http.StatusNoContent)
}
