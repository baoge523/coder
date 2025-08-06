package zap_log

import (
	"go.uber.org/zap"
	"testing"
)

func TestZapLogger(t *testing.T) {
	logger := GetLogger()

	// 添加了field和对应的值到日志中
	logger.Info("hello", "world", zap.String("namespace", "andy"), zap.String("status", "ok"))
	logger.Info("aaaa")
	logger.Infof("start %d", 11)
	logger.WithField(zap.String("namespace", "andy")).Infof("start %d", 22)
	logger.Infof("start %d", 33)
}
