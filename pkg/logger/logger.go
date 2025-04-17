package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Global logger instance
var log *zap.Logger

// Init initializes the logger with proper configuration
func Init(environment string) {
	var config zap.Config

	if environment == "production" {
		config = zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var err error
	log, err = config.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
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
	return log.Sync()
}
