package root

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go_ex01/pkg/api/handler"
	"go_ex01/pkg/api/service"
	"go_ex01/pkg/config"
)

type Server struct {
	echo   *echo.Echo
	config *config.Config
}

func New(cfg *config.Config) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	return &Server{
		echo:   e,
		config: cfg,
	}
}

func (s *Server) Setup() {
	// Setup middleware
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.CORS())

	// Setup routes
	s.setupRoutes()

	slog.Info("Server setup completed")
}

func (s *Server) setupRoutes() {
	// Health check
	s.echo.GET("/health", s.healthCheck)

	// API routes
	api := s.echo.Group("/api/v1")

	// User service and handler
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)
	userHandler.RegisterRoutes(api)

	slog.Info("Routes registered")
}

func (s *Server) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now(),
		"version":   "1.0.0",
		"service":   "go-rest-example",
	})
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.config.Server.Host, s.config.Server.Port)
	slog.Info("Starting root", "address", addr)

	if err := s.echo.Start(addr); err != nil && err != http.ErrServerClosed {
		slog.Error("Server failed to start", "error", err)
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	slog.Info("Shutting down root...")

	if err := s.echo.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
		return err
	}

	slog.Info("Server shutdown completed")
	return nil
}
