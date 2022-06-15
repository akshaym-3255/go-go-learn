package logger

import (
	nativeLogger "log"

	"go.uber.org/zap"
)

var log *zap.SugaredLogger

// InitiateLogger initiates the logger with configured values
func InitiateLogger(logLevel string, logFile string, stage string) {
	isDev := stage == "dev"

	nativeLogger.Println("\t\tINFO \tconfiguring logging with level", logLevel)
	log = Init(logLevel, logFile, isDev)
	print("\nlogger is ", &log)

}

// Debug for log.Debug
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Debugf for log.Debugf
func Debugf(template string, args ...interface{}) {
	print("\nlogger is ", &log)
	log.Debugf(template, args...)
}

// Debugw for log.Debugw
func Debugw(msg string, keysAndValues ...interface{}) {
	log.Debugw(msg, keysAndValues...)
}

// Errorf for log.Errorf
func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

// Error for log.Error
func Error(args ...interface{}) {
	log.Error(args...)
}

// Warn for log.Warn
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Warnf for log.Warnf
func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

//Warnw for log.Warnw
func Warnw(msg string, keysAndValues ...interface{}) {
	log.Warnw(msg, keysAndValues...)
}

// Info for log.Info
func Info(args ...interface{}) {
	log.Info(args...)
}

// Infof for log.Infof
func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

//Infow for log.Infow
func Infow(msg string, keysAndValues ...interface{}) {
	log.Infow(msg, keysAndValues...)
}

// Panicf for log.Panicf
func Panicf(template string, args ...interface{}) {
	log.Panicf(template, args)
}

// Panic for log.Panic
func Panic(args ...interface{}) {
	log.Panic(args)
}

// Fatal for log.Fatal
func Fatal(args ...interface{}) {
	log.Fatal(args)
}

// Fatalf for log.Fatalf
func Fatalf(template string, args ...interface{}) {
	log.Fatalf(template, args)
}
