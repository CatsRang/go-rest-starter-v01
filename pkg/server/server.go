package server

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
	logger *slog.Logger
}

func New(cfg *config.Config) *Server {
	logger := slog.With("component", "server")
	
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	
	return &Server{
		echo:   e,
		config: cfg,
		logger: logger,
	}
}

func (s *Server) Setup() {
	// Setup middleware
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.CORS())
	
	// Setup routes
	s.setupRoutes()
	
	s.logger.Info("Server setup completed")
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
	
	s.logger.Info("Routes registered")
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
	s.logger.Info("Starting server", "address", addr)
	
	if err := s.echo.Start(addr); err != nil && err != http.ErrServerClosed {
		s.logger.Error("Server failed to start", "error", err)
		return err
	}
	
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down server...")
	
	if err := s.echo.Shutdown(ctx); err != nil {
		s.logger.Error("Server forced to shutdown", "error", err)
		return err
	}
	
	s.logger.Info("Server shutdown completed")
	return nil
}