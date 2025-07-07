package receiver

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type ReceiveData struct {
	Namespace string
	ViewName  string
	// AppID 应用id，用于标识某个特定应用
	AppID string
	// OwnerUserID 用于标识主用户id
	OwnerUserID string
	// SubUserID 用于标识子用户id
	SubUserID string
	// Dimensions 用于描述产品实例的元信息
	Dimensions []map[string]interface{}
	// Metrics 用于描述产品实例的当前上报指标信息
	Metrics   []map[string]interface{}
	Timestamp int64
}

func Receive(ctx *gin.Context) {
	var receiveData ReceiveData
	err := ctx.ShouldBindJSON(&receiveData)
	if err != nil {
		handleErr(ctx, "request data parse error", err)
		return
	}
	dataByte, err := json.Marshal(receiveData)
	if err != nil {
		handleErr(ctx, "request data marshal error", err)
		return
	}
	fmt.Printf("request data %s \n", string(dataByte))

	// todo 业务功能能处理

	ctx.JSON(200, gin.H{
		"status": "success",
	})

}

func handleErr(ctx *gin.Context, title string, err error) {
	ctx.JSON(500, gin.H{
		"request_err": title + " : " + err.Error(),
	})
}

// curl http://127.0.0.1:8081/v1/test
func Test(ctx *gin.Context) {
	intn := rand.Intn(1000)
	time.Sleep(time.Duration(intn) * time.Millisecond)
	ctx.JSON(200, gin.H{
		"status": "success",
	})
}
