package template_handle

import (
	"log"
	"os"
	"strings"
	"testing"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
	Age      int
	User     User
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
	// text := `{{printf "%q" "output"}}` // 输出output
	text := `{{printf "%q" .}}` // 输出上下文的 aaaa
	masterTmpl, err := template.New("master").Parse(text)
	if err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, "aaaa"); err != nil {
		log.Fatal(err)
	}
}
