package custom_formatter

import (
	"fmt"
	"testing"
)

type User struct {
	Username string
	Password TextHidden
	Phone    TextHidden
}

type HiddenHandle interface {
	hidden(string) string
}

type PasswordHiddenHandle struct {
}

func (p PasswordHiddenHandle) hidden(source string) string {
	return "******"
}

type PhoneHiddenHandle struct {
}

func (p PhoneHiddenHandle) hidden(source string) string {
	if len(source) != 11 {
		return "***********"
	}
	return fmt.Sprintf("%s****%s", source[0:3], source[7:11])
}

type TextHidden struct {
	Value        string
	HiddenHandle HiddenHandle
}

func (p TextHidden) Format(f fmt.State, verb rune) {
	hiddenText := p.HiddenHandle.hidden(p.Value)
	_, err := f.Write([]byte(hiddenText))
	if err != nil {
		fmt.Printf("TextHidden Format error %v", err)
	}
}

func TestFormatter(t *testing.T) {
	u := User{
		Username: "aaaa",
		Password: TextHidden{Value: "123456", HiddenHandle: PasswordHiddenHandle{}},
		Phone: TextHidden{
			Value:        "17723625343",
			HiddenHandle: PhoneHiddenHandle{},
		},
	}
	fmt.Printf("%+v \n", u)
}
