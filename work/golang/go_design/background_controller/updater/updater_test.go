package updater

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {

	updater := &Updater{
		Time: time.Second,
		ValueSetter: func(ctx context.Context) (interface{}, error) {
			// 获取数据
			return "", nil
		},
		ErrorHandler: func(ctx context.Context, err error) {
			// 处理异常信息，比如上报异常指标、日志输出等
		},
	}

	if err := updater.Start(context.Background()); err != nil {
		fmt.Println("err")
		return
	}
	defer updater.Stop()

	// do other

}
