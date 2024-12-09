package logs

import "go.uber.org/zap"

// Create a logger instance
var logger, _ = zap.NewProduction()

// Info logs a message with info level
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Error logs a message with error level
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}
