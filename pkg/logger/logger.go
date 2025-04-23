package logger

import (
	appConf "github.com/aprilboiz/flight-management/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

// Global logger instance
var log *zap.Logger

// Init initializes the logger with proper configuration
func Init(environment string) {
	var zapConfig zap.Config

	// Get the application config
	appConfig := appConf.GetConfig()
	if appConfig == nil {
		panic("Failed to initialize logger: application config is nil")
	}

	// Configure based on environment
	if environment == "production" {
		zapConfig = zap.NewProductionConfig()
		zapConfig.EncoderConfig.TimeKey = "timestamp"
		zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Ensure the log directory exists
	logPath := appConfig.Logging.OutputPath
	logDir := filepath.Dir(logPath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic("Failed to create log directory: " + err.Error())
	}

	// Set output paths
	zapConfig.OutputPaths = []string{"stdout", logPath}

	// Build the logger
	var err error
	log, err = zapConfig.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	// Replace globals only after successful build
	zap.ReplaceGlobals(log)
}

// Get returns the global logger instance
func Get() *zap.Logger {
	return log
}

// Named creates a named logger
func Named(name string) *zap.Logger {
	return log.Named(name)
}

// With creates a child logger with fields
func With(fields ...zapcore.Field) *zap.Logger {
	return log.With(fields...)
}

// Sync flushes any buffered log entries
func Sync() error {
	if log != nil {
		return log.Sync()
	}
	return nil
}
