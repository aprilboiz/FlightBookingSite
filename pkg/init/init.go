package init

import (
	"os"
	"path/filepath"
	"time"

	"github.com/aprilboiz/flight-management/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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
