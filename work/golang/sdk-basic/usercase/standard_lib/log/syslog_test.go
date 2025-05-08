package log

import (
	"log"
	"log/syslog"
	"testing"
)

func TestSyslog(t *testing.T) {
	// logTypes := []string{"unixgram", "unix"}
	//	logPaths := []string{"/dev/log", "/var/run/syslog", "/var/run/log"}
	// 创建一个 syslog 服务器,默认会去连接unix的"/dev/log"等套接字程序，并将日志输出到该服务中，可以在/var/log/messages 或者/var/log/syslog看到日志
	syslogWriter, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, "my_log:")
	if err != nil {
		log.Fatal(err)
	}
	// 使用 syslog 记录日志
	log.SetOutput(syslogWriter)
	// 示例日志
	log.Println("This is a test log entry.")

}

func TestSysLog2(t *testing.T) {

	write, _ := syslog.Dial("tcp", "127.0.0.1:8898", syslog.LOG_INFO, "my_log:")
	defer write.Close()
	write.Info("this is info log")

}
