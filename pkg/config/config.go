package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

// Config struct mirrors the structure of your YAML file.
// Fields must be exported (start with uppercase) to be accessible
// outside this package.
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

// --- Singleton Implementation ---

var (
	cfg  *Config   // Private variable to hold the single instance
	once sync.Once // Ensures initialization code runs only once
)

const (
	// DefaultConfigPath defines the default location of the config file.
	// You could make this configurable via env vars or flags if needed.
	DefaultConfigPath = "pkg/config/config.yml"

	EnvironmentProduction  = "production"
	EnvironmentDevelopment = "development"
)

// loadConfig reads the YAML file and unmarshals it into the cfg variable.
// It's intended to be called only once by sync.Once.
func loadConfig() {
	filePath := DefaultConfigPath // Use the constant or get from env/flag

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		// Use Fatalf for critical config loading exceptions - app likely can't run.
		log.Fatalf("CRITICAL: Error reading config file '%s': %v", filePath, err)
	}

	var loadedConfig Config // Temporary variable to unmarshal into
	err = yaml.Unmarshal(yamlFile, &loadedConfig)
	if err != nil {
		log.Fatalf("CRITICAL: Error unmarshalling config YAML from '%s': %v", filePath, err)
	}

	// Assign to the package-level variable *after* successful loading
	cfg = &loadedConfig
	log.Println("Configuration loaded successfully from", filePath)
}

// GetConfig returns the singleton instance of the application configuration.
// It ensures that the configuration is loaded only once.
func GetConfig() *Config {
	// once.Do calls loadConfig() only on the *first* call to GetConfig().
	// Subsequent calls will skip loadConfig() but still return the cfg instance.
	once.Do(loadConfig)

	// If loadConfig panicked (due to Fatalf), the program would have already exited.
	// If loadConfig potentially returned an error instead of panicking, you'd need
	// additional checks here to ensure cfg is not nil.
	if cfg == nil {
		// This state should ideally be unreachable if loadConfig panics on failure.
		log.Panicln("CRITICAL: Configuration accessed before successful loading or loading failed without panic.")
	}
	return cfg
}

// --- Optional: Convenience Getters (Example) ---
// You can add functions to directly access specific parts of the config

func GetServerPort() int {
	return GetConfig().Server.Port
}

func GetDatabaseName() string {
	return GetConfig().Database.Name
}

func GetDatabaseConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s TimeZone=%s",
		GetConfig().Server.Host,
		GetConfig().Database.Port,
		GetConfig().Database.User,
		GetConfig().Database.Name,
		GetConfig().Database.Password,
		GetConfig().Database.Timezone,
	)
}
