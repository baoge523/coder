package zap_log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var loggerProvider = provider{logger: delegationLogger{Logger: &LoggerImpl{zap: zap.Must(DefaultConfig.Build())}}}

const timeLayout = "2006.01.02 15:04:05"

var DefaultConfig = zap.Config{
	Encoding:    "json",
	Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel), // 输出级别
	OutputPaths: []string{"stdout"},                      // 输出目的地
	EncoderConfig: zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		CallerKey:    "Caller",
		FunctionKey:  "Func",
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		TimeKey:      "time",
		EncodeTime:   zapcore.TimeEncoderOfLayout(timeLayout),
		EncodeCaller: zapcore.ShortCallerEncoder,
		LineEnding:   zapcore.DefaultLineEnding,
	},
	InitialFields: map[string]interface{}{"logger": "default"},
}

type provider struct {
	logger Logger
}

func GetLogger() Logger {
	return loggerProvider.logger
}

type delegationLogger struct {
	Logger
}

type LoggerImpl struct {
	once sync.Once
	zap  *zap.Logger
}

func (l *LoggerImpl) init() {
	l.once.Do(func() {
		l.zap = zap.Must(DefaultConfig.Build())
	})
}

func (l *LoggerImpl) Info(args ...interface{}) {
	params, fields := pickZapFields(args)
	l.zap.Info(getLogMsg(params), fields...)
}
func (l *LoggerImpl) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.zap.Info(msg)
}

// WithField 会创建一个新的logger对象
func (l *LoggerImpl) WithField(fields ...zap.Field) Logger {
	newLogger := l.zap.With(fields...)
	return &LoggerImpl{zap: newLogger}
}

type Logger interface {
	Info(args ...interface{})
	Infof(string, ...interface{})
	// 其他的就算了
	WithField(fields ...zap.Field) Logger
}

func pickZapFields(args []interface{}) ([]interface{}, []zapcore.Field) {
	var fields []zapcore.Field
	var size int
	for idx, arg := range args {
		if field, ok := arg.(zapcore.Field); ok {
			fields = append(fields, field)
			continue
		}
		if size != idx { // 如果参数中，入参和zapcore.Field存在交叉，那么就把参数按照顺序向前排序
			args[size] = arg
		}
		size++
	}
	for i := size; i < len(args); i++ {
		args[i] = nil
	}
	return args[:size], fields
}

func getLogMsg(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	msg = msg[1 : len(msg)-2] // 去掉数组的信息
	return msg
}
