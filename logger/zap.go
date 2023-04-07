package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func init() {
	// Initialize the logger configuration
	loggerConfig := zap.NewProductionConfig()

	// Set the log level
	loggerConfig.Level.SetLevel(zap.InfoLevel)

	// Set the output to a file
	logFile, _ := os.OpenFile("beprayed.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer logFile.Close()

	loggerConfig.OutputPaths = []string{
		"stdout",       // log to the console
		logFile.Name(), // log to the file
	}

	// Skip 1 caller so that the logger doesn't show itself as the caller
	loggerConfig.EncoderConfig.CallerKey = "caller"
	loggerConfig.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// Create the logger instance
	logger, _ := loggerConfig.Build()

	// Skip 1 callers to show the actual place where the log entry was generated
	logger = logger.WithOptions(
		zap.AddCallerSkip(1),
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)

	// Defer a function to flush the logger's buffer before exiting
	defer logger.Sync()

	// Create a sugared logger instance for structured logging
	sugarLogger = logger.Sugar()
}

func Info(msg string, fields ...interface{}) {
	sugarLogger.Infow(msg, fields...)
}

func Warn(msg string, fields ...interface{}) {
	sugarLogger.Warnw(msg, fields...)
}

func Error(msg string, fields ...interface{}) {
	sugarLogger.Errorw(msg, fields...)
}

func Fatal(msg string, fields ...interface{}) {
	sugarLogger.Fatalw(msg, fields...)
}
