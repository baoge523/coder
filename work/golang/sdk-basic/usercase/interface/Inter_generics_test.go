package _interface

import (
	"fmt"
	"testing"
)

// golang中的泛型
/**
1、golang1.18支持泛型(generics): https://go.dev/doc/go1.18
2、详情文档: https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md
泛型分为： 泛型方法、泛型类型    要么基于any 要么基于interface泛型
		泛型方法： func aaa[T any]() {}
		泛型类型:    type bbb[T any] struct {}   or  type bbb[T interfaceType] struct {}


*/

type MsgType interface {
	GetMsg() any
}
type SubMsgType struct {
	Code string
}

func (s *SubMsgType) GetMsg() any {
	return s.Code
}

type Sender[MT MsgType] interface {
	SendMsg(mt MT)
}

type EmailSender[MT MsgType] struct {
}

func (s *EmailSender[MT]) SendMsg(msg MT) {
	fmt.Println(msg.GetMsg())
}

func TestName(t *testing.T) {
	var send Sender[MsgType]
	send = &EmailSender[MsgType]{}
	send.SendMsg(&SubMsgType{
		Code: "aabbbcc",
	})

}
