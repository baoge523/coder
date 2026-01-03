package log

const (
	LoggerDefaultKey = "logger_in_ctx"
)

type Field struct {
	Key   string
	Value interface{}
}

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})

	With(fields ...Field) Logger

	// Sync write buffer to disk when application exit
	Sync() error
}
