package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// Configure Lumberjack for log rotation
	logWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/var/log/app/app.log", // Path to log file
		MaxSize:    100,                    // Max megabytes before log is rotated
		MaxBackups: 3,                      // Max number of old log files to keep
		MaxAge:     28,                     // Max number of days to retain old logs
		Compress:   true,                   // Compress the log files
	})

	config := zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "severity",
		MessageKey:    "message",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}

	// Use JSON encoder for structured logging
	encoder := zapcore.NewJSONEncoder(config)

	// Create a core with Lumberjack writer
	core := zapcore.NewCore(encoder, logWriter, zap.InfoLevel)

	// Combine core with additional outputs (stdout)
	stdoutSync := zapcore.AddSync(zapcore.Lock(os.Stdout))
	combinedCore := zapcore.NewTee(
		core, // Logs to file with rotation
		zapcore.NewCore(encoder, stdoutSync, zap.InfoLevel), // Logs to stdout
	)

	// Build the logger
	Logger := zap.New(combinedCore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	defer Logger.Sync()

	// Test the logger
	Logger.Info("Logger initialized successfully")
	Logger.Error("This is an error log")
}
