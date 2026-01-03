package log

import (
	"context"
)

func Debugf(format string, args ...interface{}) {
	GetDefaultLogger().Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	GetDefaultLogger().Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	GetDefaultLogger().Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	GetDefaultLogger().Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	GetDefaultLogger().Fatalf(format, args...)
}

// log with context

func DebugContextf(ctx context.Context, format string, args ...interface{}) {
	if logger, ok := ctx.Value(LoggerDefaultKey).(Logger); ok {
		logger.Debugf(format, args...)
		return
	}
	GetDefaultLogger().Debugf(format, args...)
}

func InfoContextf(ctx context.Context, format string, args ...interface{}) {
	if logger, ok := ctx.Value(LoggerDefaultKey).(Logger); ok {
		logger.Infof(format, args...)
		return
	}
	GetDefaultLogger().Infof(format, args...)
}

func WarnContextf(ctx context.Context, format string, args ...interface{}) {
	if logger, ok := ctx.Value(LoggerDefaultKey).(Logger); ok {
		logger.Warnf(format, args...)
		return
	}
	GetDefaultLogger().Warnf(format, args...)
}

func ErrorContextf(ctx context.Context, format string, args ...interface{}) {
	if logger, ok := ctx.Value(LoggerDefaultKey).(Logger); ok {
		logger.Errorf(format, args...)
		return
	}
	GetDefaultLogger().Errorf(format, args...)
}

func FatalContextf(ctx context.Context, format string, args ...interface{}) {
	if logger, ok := ctx.Value(LoggerDefaultKey).(Logger); ok {
		logger.Fatalf(format, args...)
		return
	}
	GetDefaultLogger().Fatalf(format, args...)
}

func Printf(format string, args ...interface{}) {
	GetDefaultLogger().Infof(format, args...)
}

// WithFields return new logger that with addition fields
func WithFields(ctx context.Context, fields ...Field) Logger {
	if logger, ok := ctx.Value(LoggerDefaultKey).(Logger); ok {
		return logger.With(fields...)
	}
	return GetDefaultLogger().With(fields...)
}

// WithContext store logger and return child context
func WithContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, LoggerDefaultKey, logger)
}

func NewLogContext(ctx context.Context, field ...Field) (context.Context, Logger) {
	if len(field) != 0 {
		logger := GetDefaultLogger().With(field...)
		return context.WithValue(ctx, LoggerDefaultKey, logger), logger
	}
	logger := GetDefaultLogger().With(field...)
	return context.WithValue(ctx, LoggerDefaultKey, logger), logger
}
func Sync() error {
	return GetDefaultLogger().Sync()
}
