package log

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	defaultLogger Logger

	mu sync.RWMutex
)

const timeLayout = "2006.01.02 15:04:05"

// lumberjack 实现文件存储日志文件，并按天和文件大小滚动
func initRollFileLogger() io.Writer {
	// 按天生成文件名
	today := time.Now().Format("2006-01-02")
	//
	filename := fmt.Sprintf("./grpc_demo/hello_world/server/logs/app_%s.log", today)

	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    100, // MB
		MaxBackups: 7,   // 保留7天的日志
		MaxAge:     7,   // 保留7天
		Compress:   true,
		LocalTime:  true,
	}
	return lumberjackLogger
}

func initConsoleLogger() io.Writer {
	return os.Stdout
}

func init() {
	config := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "Caller",
		FunctionKey:    "Func",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 日志级别信息 -- 可以选择大小写 或者是否带有颜色
		SkipLineEnding: false,                       // 为true时，设置的行结尾无效
		TimeKey:        "time",
		EncodeTime:     zapcore.TimeEncoderOfLayout(timeLayout), //  时间格式化器
		EncodeCaller:   zapcore.ShortCallerEncoder,              // short caller info -- 这里是调用的函数
		LineEnding:     zapcore.DefaultLineEnding,               // 行结尾，换行
	}

	encoder := zapcore.NewJSONEncoder(config)

	// 本质也是json，但是元信息不按照json输出
	// 2026.01.02 19:45:15	INFO	Received: Alice
	// {"caller_ip": "[::1]:63656", "caller_service": "hello_world_caller", "caller_version": "1.0.0", "callee_service": "hello_world.Greeter", "callee_method": "SayHello"}
	//encoder := zapcore.NewConsoleEncoder(config)

	// todo write log to file and console or other (kafka ...)
	console := zapcore.NewCore(encoder, zapcore.AddSync(initConsoleLogger()), zapcore.InfoLevel)
	rollFile := zapcore.NewCore(encoder, zapcore.AddSync(initRollFileLogger()), zapcore.InfoLevel)

	core := zapcore.NewTee(console, rollFile)
	zapLog := zap.New(core)
	defaultLogger = &zapLogger{
		logger: zapLog,
	}

}

type zapLogger struct {
	logger *zap.Logger
}

func GetDefaultLogger() Logger {
	mu.RLock()
	defer mu.RUnlock()
	return defaultLogger
}

func (dl *zapLogger) Debugf(format string, args ...interface{}) {
	dl.logger.Debug(fmt.Sprintf(format, args...))
}

func (dl *zapLogger) Infof(format string, args ...interface{}) {
	dl.logger.Info(fmt.Sprintf(format, args...))
}
func (dl *zapLogger) Warnf(format string, args ...interface{}) {
	dl.logger.Warn(fmt.Sprintf(format, args...))
}
func (dl *zapLogger) Errorf(format string, args ...interface{}) {
	dl.logger.Error(fmt.Sprintf(format, args...))
}

func (dl *zapLogger) Fatalf(format string, args ...interface{}) {
	dl.logger.Fatal(fmt.Sprintf(format, args...))
}
func (dl *zapLogger) With(fields ...Field) Logger {
	if len(fields) == 0 {
		return dl
	}

	zapFields := make([]zap.Field, len(fields))
	for index, field := range fields {
		zapFields[index] = zap.Any(field.Key, field.Value)
	}

	newLogger := dl.logger.With(zapFields...)
	return &zapLogger{
		logger: newLogger,
	}
}

func (dl *zapLogger) Sync() error {
	return dl.logger.Sync()
}
