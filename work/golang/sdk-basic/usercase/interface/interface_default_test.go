package _interface

import "testing"




type AFace interface {
	a() string
}

var aa AFace

type AImpl struct {

}

func (a *AImpl) a() string {
	return "aa"
}


func TestInterfaceDefault(test *testing.T) {
	aa.a()
}