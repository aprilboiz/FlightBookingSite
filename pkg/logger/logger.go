package logger

import (
	"os"
	"path/filepath"
	"time"

	"github.com/aprilboiz/flight-management/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Global logger instance
var log *zap.Logger

// InitLogger initializes the logger with proper configuration
func InitLogger(environment string) *zap.Logger {
	var zapConfig zap.Config

	// Configure based on environment
	if environment == "production" {
		zapConfig = zap.NewProductionConfig()
		zapConfig.EncoderConfig.TimeKey = "timestamp"
		zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		zapConfig.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	} else {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapConfig.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		}
	}

	// Set log level from config
	level := zapcore.InfoLevel
	if err := level.UnmarshalText([]byte(config.GetConfig().Logging.Level)); err != nil {
		level = zapcore.InfoLevel
	}
	zapConfig.Level = zap.NewAtomicLevelAt(level)

	// Ensure the log directory exists
	logPath := config.GetConfig().Logging.OutputPath
	logDir := filepath.Dir(logPath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic("Failed to create log directory: " + err.Error())
	}

	// Set output paths
	zapConfig.OutputPaths = []string{"stdout", logPath}
	zapConfig.ErrorOutputPaths = []string{"stderr", logPath}

	// Add caller information
	zapConfig.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// Build the logger
	log, err := zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	// Replace globals
	zap.ReplaceGlobals(log)
	return log
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

// Debug logs a debug message
func Debug(msg string, fields ...zapcore.Field) {
	log.Debug(msg, fields...)
}

// Info logs an info message
func Info(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}

// Warn logs a warning message
func Warn(msg string, fields ...zapcore.Field) {
	log.Warn(msg, fields...)
}

// Error logs an error message
func Error(msg string, fields ...zapcore.Field) {
	log.Error(msg, fields...)
}

// Fatal logs a fatal message and then calls os.Exit(1)
func Fatal(msg string, fields ...zapcore.Field) {
	log.Fatal(msg, fields...)
}

// WithError adds an error field to the logger
func WithError(err error) *zap.Logger {
	return log.With(zap.Error(err))
}

// WithRequestID adds a request ID field to the logger
func WithRequestID(requestID string) *zap.Logger {
	return log.With(zap.String("request_id", requestID))
}

// WithUserID adds a user ID field to the logger
func WithUserID(userID uint) *zap.Logger {
	return log.With(zap.Uint("user_id", userID))
}

// WithDuration adds a duration field to the logger
func WithDuration(duration time.Duration) *zap.Logger {
	return log.With(zap.Duration("duration", duration))
}

// WithString adds a string field to the logger
func WithString(key, value string) *zap.Logger {
	return log.With(zap.String(key, value))
}

// WithInt adds an integer field to the logger
func WithInt(key string, value int) *zap.Logger {
	return log.With(zap.Int(key, value))
}

// WithUint adds an unsigned integer field to the logger
func WithUint(key string, value uint) *zap.Logger {
	return log.With(zap.Uint(key, value))
}

// WithFloat64 adds a float64 field to the logger
func WithFloat64(key string, value float64) *zap.Logger {
	return log.With(zap.Float64(key, value))
}

// WithBool adds a boolean field to the logger
func WithBool(key string, value bool) *zap.Logger {
	return log.With(zap.Bool(key, value))
}

// WithTime adds a time field to the logger
func WithTime(key string, value time.Time) *zap.Logger {
	return log.With(zap.Time(key, value))
}

// WithAny adds any field to the logger
func WithAny(key string, value any) *zap.Logger {
	return log.With(zap.Any(key, value))
}
