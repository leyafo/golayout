package logger

import "go.uber.org/zap"

// Debugf is the wrapper of default logger Debugf.
func Debugf(template string, args ...interface{}) {
	sugaredLogger.Debugf(template, args...)
}

// Infof is the wrapper of default logger Infof.
func Infof(template string, args ...interface{}) {
	sugaredLogger.Infof(template, args...)
}

// Warnf is the wrapper of default logger Warnf.
func Warnf(template string, args ...interface{}) {
	sugaredLogger.Warnf(template, args...)
}

// Errorf is the wrapper of default logger Errorf.
func Errorf(template string, args ...interface{}) {
	sugaredLogger.Errorf(template, args...)
}

// Debug is the wrapper of default logger Debug.
func Debug(args ...interface{}) {
	sugaredLogger.Debug(args...)
}

// Info is the wrapper of default logger Info.
func Info(args ...interface{}) {
	sugaredLogger.Info(args...)
}

// Warn is the wrapper of default logger Warn.
func Warn(args ...interface{}) {
	sugaredLogger.Warn(args...)
}

// Error is the wrapper of default logger Error.
func Error(args ...interface{}) {
	sugaredLogger.Error(args...)
}

// Fatal is the wrapper of default logger Fatal.
func Fatal(args ...interface{}) {
	sugaredLogger.Fatal(args...)
}

// Fatalf is the wrapper of default logger Fatalf.
func Fatalf(template string, args ...interface{}) {
	sugaredLogger.Fatalf(template, args...)
}

func FastLogger() *zap.Logger {
	return zapLogger
}

// Sync syncs all logs, must be called after calling Init().
func Sync() {
	sugaredLogger.Sync()
	zapLogger.Sync()
}
