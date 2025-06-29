package config

import (
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Environment string         `yaml:"environment"`
	Server      ServerConfig   `yaml:"server"`
	Database    DatabaseConfig `yaml:"database"`
	Logging     LoggingConfig  `yaml:"logging"`
}

type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type DatabaseConfig struct {
	Port     int                `yaml:"port"`
	Host     string             `yaml:"host"`
	Type     string             `yaml:"type"`
	User     string             `yaml:"user"`
	Password string             `yaml:"password"`
	Name     string             `yaml:"name"`
	Timezone string             `yaml:"timezone"`
	Options  map[string]int     `yaml:"options"`
	Init     DatabaseInitConfig `yaml:"init"`
}

type DatabaseInitConfig struct {
	SeedPath string `yaml:"seed_path"`
	RunSeed  bool   `yaml:"run_seed"`
}

type LoggingConfig struct {
	Level      string `yaml:"level"`
	Format     string `yaml:"format"`
	OutputPath string `yaml:"output_path"`
}

var (
	cfg  *Config   // Private variable to hold the single instance
	once sync.Once // Ensures initialization code runs only once
)

const (
	DefaultConfigPath = "pkg/config/config.yml"

	EnvironmentProduction  = "production"
	EnvironmentDevelopment = "development"
)

// loadConfig reads the YAML file and unmarshals it into the cfg variable.
// It's intended to be called only once by sync.Once.
func loadConfig() {
	filePath := DefaultConfigPath // Use the constant or get from env/flag
	log := zap.L()

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		// Use Fatal for critical config loading exceptions - app likely can't run.
		log.Fatal("Failed to read config file",
			zap.String("path", filePath),
			zap.Error(err))
	}

	var loadedConfig Config // Temporary variable to unmarshal into
	err = yaml.Unmarshal(yamlFile, &loadedConfig)
	if err != nil {
		log.Fatal("Failed to unmarshal config YAML",
			zap.String("path", filePath),
			zap.Error(err))
	}

	// Assign to the package-level variable *after* successful loading
	cfg = &loadedConfig
	log.Info("Configuration loaded successfully",
		zap.String("path", filePath))
}

// GetConfig returns the singleton instance of the application configuration.
// It ensures that the configuration is loaded only once.
func GetConfig() *Config {
	// once.Do calls loadConfig() only on the *first* call to GetConfig().
	// Subsequent calls will skip loadConfig() but still return the cfg instance.
	once.Do(loadConfig)

	if cfg == nil {
		zap.L().Panic("Configuration accessed before successful loading or loading failed without panic")
	}
	return cfg
}

func GetDatabaseConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s TimeZone=%s",
		GetConfig().Database.Host,
		GetConfig().Database.Port,
		GetConfig().Database.User,
		GetConfig().Database.Name,
		GetConfig().Database.Password,
		GetConfig().Database.Timezone,
	)
}
