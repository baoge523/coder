package custom_formatter

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	u := User{
		Username: "aaaa",
		Password: TextHidden{Value: "123456", HiddenHandle: PasswordHiddenHandle{}},
		Phone: TextHidden{
			Value:        "17723625343",
			HiddenHandle: PhoneHiddenHandle{},
		},
	}
	marshal, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
}

func (p TextHidden) MarshalJSON() ([]byte, error) {
	hiddenText := p.HiddenHandle.hidden(p.Value)
	// 注意： 这里需要双引号
	return []byte(fmt.Sprintf("%q", hiddenText)), nil
}
