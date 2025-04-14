package template_handle

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"text/template"
)

type TUser struct {
	Name     string
	BirthDay string
	High     float32
	Hobby    struct {
		Name     string
		Deep     int
		HowOften string
	}
}

func TestJson(t *testing.T) {

	u := TUser{
		Name:     "AAA",
		BirthDay: "1111-11-11",
		High:     180.4,
		Hobby: struct {
			Name     string
			Deep     int
			HowOften string
		}{Name: "football", Deep: 5, HowOften: "sometimes"},
	}

	funcMap := template.FuncMap{
		"json": json.Marshal,
	}

	ub, _ := json.Marshal(u)

	fmt.Println(string(ub))

	// json 拼接模版
	// constContent := `{"title":"go template","name":"{{.Name}}","birth":"{{.BirthDay}}","high":{{.High}},"hobby":"{{.Hobby.Name}}-{{.Hobby.Deep}}-{{.Hobby.HowOften}}"}`

	// 输出json模版
	constContent := `{"title":"go template","name":"{{.Name}}","birth":"{{.BirthDay}}","high":{{.High}},"hobby": {{.Hobby | json | printf "%s" }}}`

	temp, _ := template.New("test").Funcs(funcMap).Parse(constContent)

	_ = temp.Execute(os.Stdout, &u)

	fmt.Println()
	fmt.Println("success")

}
