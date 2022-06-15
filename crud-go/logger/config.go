package logger

import (
	"fmt"
	"net/url"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

const maxSize = 10 //MB
const maxBackups = 30
const maxDays = 90

const (
	// DebugLevelStr DPanic, Panic and Fatal level can not be set by user
	DebugLevelStr string = "debug"
	// InfoLevelStr for setting info level
	InfoLevelStr string = "info"
	// WarningLevelStr for setting warning level
	WarningLevelStr string = "warning"
	// ErrorLevelStr for setting error level
	ErrorLevelStr string = "error"
)

var (
	// Logger to expose logger as global var
	Logger zap.Logger
)

type lumberjackSink struct {
	*lumberjack.Logger
}

// Sync to manage log rotation
func (lumberjackSink) Sync() error {
	return nil
}

// Init init the logger
func Init(logLevel string, logFile string, isDev bool) *zap.SugaredLogger {
	var level zapcore.Level

	switch logLevel {
	case DebugLevelStr:
		level = zap.DebugLevel

	case InfoLevelStr:
		level = zap.InfoLevel

	case WarningLevelStr:
		level = zap.WarnLevel

	case ErrorLevelStr:
		level = zap.ErrorLevel

	default:
		return nil
	}

	encoderConfig := GetEncoding()

	ll := lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    maxSize, //MB
		MaxBackups: maxBackups,
		MaxAge:     maxDays, //days
		Compress:   true,
	}
	_ = zap.RegisterSink("lumberjack", func(*url.URL) (zap.Sink, error) {
		return lumberjackSink{
			Logger: &ll,
		}, nil
	})

	loggerConfig := zap.Config{
		Level:         zap.NewAtomicLevelAt(level),
		Development:   isDev,
		Encoding:      "json",
		EncoderConfig: encoderConfig,
		OutputPaths:   []string{fmt.Sprintf("lumberjack:%s", logFile), "stdout"},
	}

	globalLogger, err := loggerConfig.Build()

	if err != nil {
		panic(fmt.Sprintf("build zap logger from config error: %v", err))
	}

	_ = zap.ReplaceGlobals(globalLogger)
	_ = globalLogger.Sync()

	Logger := globalLogger.Sugar()
	fmt.Print("logger initialized")
	return Logger
}

// GetEncoding logger encoding
func GetEncoding() zapcore.EncoderConfig {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return encoderConfig
}
