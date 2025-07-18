package template_handle

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"text/template"
	"time"
)

type Inventory struct {
	Material string
	Count    uint
	Age      int
	User     User
	UserList []User
}

type User struct {
	Name  string
	Hobby string
	Age   int
}

// 级联对象时，通过{{.User.Name}}获取级联对象的属性信息
func TestTemplateExtend(t *testing.T) {

	sweaters := Inventory{Material: "wool", Count: 17, User: User{
		Name:  "zhangsan",
		Hobby: "basketball",
	}}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}} {{.User.Name}} {{.User.Hobby}}\n")
	if err != nil {
		panic(err)
	}

	// 将模版替换的结果输出到标准输出中
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

}

func TestTemplate(t *testing.T) {

	sweaters := Inventory{Material: "wool", Count: 17}
	tmpl, err := template.New("test").Parse("{{.Count    }} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}

	// 将模版替换的结果输出到标准输出中
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

/*
*
- `{{block "list" .}}`：定义一个名为 "list" 的块，使用当前上下文（`.`）。
- `{{"\n"}}`：插入一个换行符。
- `{{range .}}`：遍历当前上下文的每个元素。
- `{{println "-" .}}`：打印每个元素，前面加上一个破折号，并换行。
- `{{end}}`：结束 `range` 循环。
- `{{end}}`：结束 "list" 块。
*/
func TestTemplateBlock(t *testing.T) {
	const (
		master  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		overlay = `{{define "list"}} {{join . ", "}}{{end}} `
	)
	var (
		funcs     = template.FuncMap{"join": strings.Join} // 定义一个func名称为join，实现是 strings.Join
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)

	// 创建一个text/template模版master，注入func，并解析模版
	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}

}

func TestTemplatePrint(t *testing.T) {
	text := `{{printf "%q" "output"}}` // 输出output
	//text := `{{printf "%q" .}}` // 输出上下文的 aaaa
	masterTmpl, err := template.New("master").Parse(text)
	if err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, "aaaa"); err != nil {
		log.Fatal(err)
	}
}

func TestTemplateIF(t *testing.T) {

	templateStr := `{{- if eq .current_level_fmt "提示"}}
通知标题：cos带宽超限告警 -提示
{{- else if eq .current_level_fmt "严重" }}
通知标题：cos带宽超限告警 -严重
{{- else }}
通知标题：cos带宽超限告警 -紧急
{{- end}}
通知时间：{{.first_trigger_time_fmt}}

{{- if eq .current_level_fmt "提示"}}
通知级别：通知 {{.current_level_fmt}}
{{- else if eq .current_level_fmt "严重" }}
通知级别：一般告警  {{.current_level_fmt}}
{{- else }}
通知级别：紧急告警  {{.current_level_fmt}}
{{- end}}`

	data := make(map[string]string)
	data["current_level_fmt"] = "提示"
	data["first_trigger_time_fmt"] = time.Now().String()

	parse, _ := template.New("test").Parse(templateStr)
	err := parse.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("err %w", err)
	}

}

func TestFlag(t *testing.T) {
	content := `[[aaa]] = {{.}}`

	parse, _ := template.New("test").Parse(content)
	parse.Execute(os.Stdout, "hello")

}
func TestArray(t *testing.T) {
	sweaters := Inventory{Material: "wool", Count: 17, UserList: []User{
		{
			Name:  "zhangsan",
			Hobby: "basketball",
		},
	}}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of userlist[0].name =  {{ with $info := index .UserList 0}}{{$info.Name}}{{end}}\n")
	if err != nil {
		panic(err)
	}

	// 将模版替换的结果输出到标准输出中
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

func TestCondition(t *testing.T) {
	content := "{{ if eq .ContextDimension.Lang \"zh\"}} 用户ID:{{ .MonitorViewDimension.subuin }}｜服务名称:{{ .TableDimension.service_group_name }}｜服务ID:{{ .MonitorViewDimension.taskid }}｜服务描述:{{ .TableDimension.service_description }}｜服务组ID:{{ .TableDimension.service_group_id }} {{else}} SubUin:{{ .MonitorViewDimension.subuin }}｜ServiceName:{{ .TableDimension.service_group_name }}｜ServiceID:{{ .MonitorViewDimension.taskid }}｜ServiceDescribe:{{ .TableDimension.service_description }}｜ServiceGroupID:{{ .TableDimension.service_group_id }} {{end}}"

	params := map[string]map[string]string {
		"MonitorViewDimension": {"subuin":"uin1111111","taskid":"taskid3333333"},
		"TableDimension": {"service_group_name":"group_name2222222","service_description":"service_description444444444","service_group_id":"service_group_id555555"},
		"ContextDimension" :  {"Lang": "en"},
	}
	parse, err := template.New("test").Parse(content)
	if err != nil {
		fmt.Println(err)
	}
	parse.Execute(os.Stdout, params)

}

func TestMap(t *testing.T) {
	params := make(map[string]map[string]string)
	m1 := make(map[string]string)
	m1["bb"] = "bb"
	params["aa"] = m1
	fmt.Println(params)
}