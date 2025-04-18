package log

import (
	"log"
	"log/syslog"
	"testing"
)

func TestSyslog(t *testing.T) {

	//// 默认是连接的unix的默认输出
	//writer, _ := syslog.New(syslog.LOG_INFO, "my_log:")
	//defer writer.Close()
	//writer.Info("hello syslog")

	// 创建一个 syslog 服务器
	syslogWriter, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, "my_log:")
	if err != nil {
		log.Fatal(err)
	}
	// 使用 syslog 记录日志
	log.SetOutput(syslogWriter)
	// 示例日志
	log.Println("This is a test log entry.")

}
