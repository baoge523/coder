package main

import (
	"flag"
	"fmt"
)

type TheadPoolParam struct {
	ConnMinSize       int
	ConnMaxSize       int
	ConnIdleTime      int64 //连接空闲时间
	ConnWaitQueueSize int64 // 连接等待队列大小

}

func main() {
	theadPoolParam := TheadPoolParamFlagHand()
	business(theadPoolParam)
}

func TheadPoolParamFlagHand() *TheadPoolParam {

	const (
		ConnMinSizeName    = "connMinSize"
		ConnMinSizeDefault = 2
		ConnMinSizeUsage   = "connection min size"

		ConnMaxSizeName    = "connMaxSize"
		ConnMaxSizeDefault = 15
		ConnMaxSizeUsage   = "connection max size"

		ConnIdleTimeName    = "connIdleTime"
		ConnIdleTimeDefault = 60000
		ConnIdleTimeUsage   = "connection idle timeout"

		ConnWaitQueueName        = "connWaitQueue"
		ConnWaitQueueSizeDefault = 500
		ConnWaitQueueSizeUsage   = "connection wait queue size"
	)
	var ConnMinSize int
	var ConnMaxSize int
	var ConnIdleTime int64
	var ConnWaitQueueSize int64

	flag.IntVar(&ConnMinSize, ConnMinSizeName, ConnMinSizeDefault, ConnMinSizeUsage)
	flag.IntVar(&ConnMaxSize, ConnMaxSizeName, ConnMaxSizeDefault, ConnMaxSizeUsage)
	flag.Int64Var(&ConnIdleTime, ConnIdleTimeName, ConnIdleTimeDefault, ConnIdleTimeUsage)
	flag.Int64Var(&ConnWaitQueueSize, ConnWaitQueueName, ConnWaitQueueSizeDefault, ConnWaitQueueSizeUsage)

	// 让定义的flag生效 must
	flag.Parse()

	return &TheadPoolParam{
		ConnMinSize:       ConnMinSize,
		ConnMaxSize:       ConnMaxSize,
		ConnIdleTime:      ConnIdleTime,
		ConnWaitQueueSize: ConnWaitQueueSize,
	}
}

// mock 业务逻辑
func business(param *TheadPoolParam) {
	fmt.Printf("business: %+v \n", param)
}
