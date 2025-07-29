package root

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"go_ex01/pkg/config"
	"go_ex01/pkg/util"
)

var configPath string

// NewRootCommand creates and returns the root cobra command
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "go_ex01-api",
		Short: "Example REST API root",
		Long:  "A simple go_ex01 demonstrating Go REST API with Echo, Viper, Cobra, and slog",
		Run:   rootRun,
	}

	rootCmd.Flags().StringVarP(&configPath, "config", "c", "./config.yaml", "Path to config file")
	cobra.OnInitialize(initConfig)

	return rootCmd
}

func initConfig() {
	cfg := config.Get()
	if err := cfg.Load(configPath); err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	util.InitLogger(cfg.Log.Level)
	slog.Info("Configuration loaded", "config_path", configPath)
	slog.Info("Config", "conf", cfg)
}

func rootRun(cmd *cobra.Command, args []string) {
	cfg := config.Get()

	srv := New(cfg)
	srv.Setup()

	go func() {
		if err := srv.Start(); err != nil {
			slog.Error("Server failed to start", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	}
	slog.Info("Server exited")
}
