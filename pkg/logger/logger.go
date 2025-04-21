package logger

import (
	"fmt"
	"log"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	instance *zap.Logger
	once     sync.Once
)

// Custom color encoder
func colorEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		var color string
		switch level {
		case zapcore.DebugLevel:
			color = "\x1b[36m" // Cyan
		case zapcore.InfoLevel:
			color = "\x1b[32m" // Green
		case zapcore.WarnLevel:
			color = "\x1b[33m" // Yellow
		case zapcore.ErrorLevel:
			color = "\x1b[31m" // Red
		case zapcore.FatalLevel:
			color = "\x1b[35m" // Magenta
		default:
			color = "\x1b[0m" // Reset
		}
		enc.AppendString(fmt.Sprintf("%s%s\x1b[0m", color, level.CapitalString()))
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// Init initializes the logger singleton
func Init(level string) *zap.Logger {
	once.Do(func() {
		// Parse the log level
		var zapLevel zapcore.Level
		if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
			zapLevel = zapcore.InfoLevel
		}

		// Create core with color encoder
		core := zapcore.NewCore(
			colorEncoder(),
			zapcore.AddSync(zapcore.Lock(os.Stdout)),
			zapLevel,
		)

		// Create logger with options
		logger := zap.New(core,
			zap.AddCaller(),
			zap.AddStacktrace(zapcore.ErrorLevel),
		)

		instance = logger
	})
	return instance
}

// Get returns the logger instance
func Get() *zap.Logger {
	if instance == nil {
		log.Fatal("Logger not initialized. Call Init() first")
	}
	return instance
}

// Info logs an info message
func Info(msg string, fields ...zap.Field) {
	Get().Info(msg, fields...)
}

// Error logs an error message
func Error(msg string, fields ...zap.Field) {
	Get().Error(msg, fields...)
}

// Debug logs a debug message
func Debug(msg string, fields ...zap.Field) {
	Get().Debug(msg, fields...)
}

// Warn logs a warning message
func Warn(msg string, fields ...zap.Field) {
	Get().Warn(msg, fields...)
}

// Fatal logs a fatal message and exits
func Fatal(msg string, fields ...zap.Field) {
	Get().Fatal(msg, fields...)
}
