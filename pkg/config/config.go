package config

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
}

var (
	instance *Config
	once     sync.Once
)

func Get() *Config {
	once.Do(func() {
		instance = &Config{}
	})
	return instance
}

func (c *Config) Load(configPath string) error {
	// ---- 10. bind config file
	{
		configDir := filepath.Dir(configPath)
		configName := strings.TrimSuffix(filepath.Base(configPath), filepath.Ext(configPath))
		configType := strings.TrimPrefix(filepath.Ext(configPath), ".")

		viper.SetConfigName(configName)
		viper.SetConfigType(configType)
		viper.AddConfigPath(configDir)
	}

	// ---- 20. bind ENV vars
	{
		// Enable automatic env var reading
		viper.AutomaticEnv()

		// Set environment variable prefix
		viper.SetEnvPrefix("GO_EX01")

		// Replace dots and dashes with underscores in env vars
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

		// Bind specific environment variables
		_ = viper.BindEnv("server.host", "GO_EX01_SERVER_HOST")
		_ = viper.BindEnv("server.port", "GO_EX01_SERVER_PORT")
		_ = viper.BindEnv("log.level", "GO_EX01_LOG_LEVEL")

		// Set default values
		viper.SetDefault("server.host", "localhost")
		viper.SetDefault("server.port", 8080)
		viper.SetDefault("log.level", "info")

		if err := viper.ReadInConfig(); err != nil {
			// Config file is optional when using env vars
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return err
			}
		}
	}

	return viper.Unmarshal(c)
}
