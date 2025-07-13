package initialize

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

type LoggerConfig struct {
	LogLevel    string
	LogFile     string
	MaxSize     int
	MaxBackups  int
	MaxAge      int
	Compress    bool
	Development bool
}

// NewLogger tạo logger với cấu hình linh hoạt
func NewLogger(config LoggerConfig) (*zap.Logger, error) {
	level, err := parseLogLevel(config.LogLevel)
	if err != nil {
		return nil, err
	}

	// Tạo thư mục log nếu cần
	if config.LogFile != "" {
		if err := os.MkdirAll(filepath.Dir(config.LogFile), 0755); err != nil {
			return nil, fmt.Errorf("failed to create log directory: %w", err)
		}
	}

	core := zapcore.NewTee(
		createConsoleCore(level, config.Development),
		createFileCore(level, config),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	
	// Thay thế global logger của Zap
	zap.ReplaceGlobals(logger)
	
	return logger, nil
}

func parseLogLevel(level string) (zapcore.Level, error) {
	switch level {
	case "debug":
		return zapcore.DebugLevel, nil
	case "info":
		return zapcore.InfoLevel, nil
	case "warn":
		return zapcore.WarnLevel, nil
	case "error":
		return zapcore.ErrorLevel, nil
	default:
		return zapcore.InfoLevel, fmt.Errorf("invalid log level: %s, using info", level)
	}
}

func createConsoleCore(level zapcore.Level, development bool) zapcore.Core {
	encoderConfig := zap.NewProductionEncoderConfig()
	if development {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.Lock(os.Stderr),
		level,
	)
}

func createFileCore(level zapcore.Level, config LoggerConfig) zapcore.Core {
	if config.LogFile == "" {
		return nil
	}

	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.LogFile,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
		LocalTime:  true,
	})

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		fileWriter,
		level,
	)
}