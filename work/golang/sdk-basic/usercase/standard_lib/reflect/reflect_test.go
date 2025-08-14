package reflect

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cloudwego/eino/components/model"
)

func TestReflect(t *testing.T) {
	InitLLM()
}

// LLMDep LLM对象集合
type LLMDep struct {
	AlarmChatModel model.ToolCallingChatModel `id:"alarm_chat_model"`
}

// InitLLM 初始化LLM依赖
func InitLLM() (*LLMDep, error) {
	dep := &LLMDep{}
	val := reflect.ValueOf(dep).Elem()
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		// field := val.Field(i)
		fieldType := typ.Field(i)
		idInTag := fieldType.Tag.Get("id") // alarm_chat_model
		fmt.Println(idInTag)
	}
	return dep, nil
}
