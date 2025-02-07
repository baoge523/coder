package template_handle

import (
	"log"
	"os"
	"testing"
	"text/template"
)

// {{/* a comment */}}
//{{- /* a comment with white space trimmed from preceding and following text */ -}}
//	A comment; discarded. May contain newlines.
//	Comments do not nest and must start and end at the
//	delimiters, as shown here.
//
//{{pipeline}}
//	The default textual representation (the same as would be
//	printed by fmt.Print) of the value of the pipeline is copied
//	to the output.
//
//{{if pipeline}} T1 {{end}}
//	If the value of the pipeline is empty, no output is generated;
//	otherwise, T1 is executed. The empty values are false, 0, any
//	nil pointer or interface value, and any array, slice, map, or
//	string of length zero.
//	Dot is unaffected.
//
//{{if pipeline}} T1 {{else}} T0 {{end}}
//	If the value of the pipeline is empty, T0 is executed;
//	otherwise, T1 is executed. Dot is unaffected.
//
//{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
//	To simplify the appearance of if-else chains, the else action
//	of an if may include another if directly; the effect is exactly
//	the same as writing
//		{{if pipeline}} T1 {{else}}{{if pipeline}} T0 {{end}}{{end}}
//
//{{range pipeline}} T1 {{end}}
//	The value of the pipeline must be an array, slice, map, or channel.
//	If the value of the pipeline has length zero, nothing is output;
//	otherwise, dot is set to the successive elements of the array,
//	slice, or map and T1 is executed. If the value is a map and the
//	keys are of basic type with a defined order, the elements will be
//	visited in sorted key order.
//
//{{range pipeline}} T1 {{else}} T0 {{end}}
//	The value of the pipeline must be an array, slice, map, or channel.
//	If the value of the pipeline has length zero, dot is unaffected and
//	T0 is executed; otherwise, dot is set to the successive elements
//	of the array, slice, or map and T1 is executed.
//
//{{break}}
//	The innermost {{range pipeline}} loop is ended early, stopping the
//	current iteration and bypassing all remaining iterations.
//
//{{continue}}
//	The current iteration of the innermost {{range pipeline}} loop is
//	stopped, and the loop starts the next iteration.
//
//{{template "name"}}
//	The template with the specified name is executed with nil data.
//
//{{template "name" pipeline}}
//	The template with the specified name is executed with dot set
//	to the value of the pipeline.
//
//{{block "name" pipeline}} T1 {{end}}
//	A block is shorthand for defining a template
//		{{define "name"}} T1 {{end}}
//	and then executing it in place
//		{{template "name" pipeline}}
//	The typical use is to define a set of root templates that are
//	then customized by redefining the block templates within.
//
//{{with pipeline}} T1 {{end}}
//	If the value of the pipeline is empty, no output is generated;
//	otherwise, dot is set to the value of the pipeline and T1 is
//	executed.
//
//{{with pipeline}} T1 {{else}} T0 {{end}}
//	If the value of the pipeline is empty, dot is unaffected and T0
//	is executed; otherwise, dot is set to the value of the pipeline
//	and T1 is executed.
//
//{{with pipeline}} T1 {{else with pipeline}} T0 {{end}}
//	To simplify the appearance of with-else chains, the else action
//	of a with may include another with directly; the effect is exactly
//	the same as writing
//		{{with pipeline}} T1 {{else}}{{with pipeline}} T0 {{end}}{{end}}

// {{"\n"}} 表示换行符
func TestTemplateActionComment(t *testing.T) {
	// {{/* a comment */}}
	user := User{Name: "lisi", Hobby: "cook"}
	tempText := `user name {{.Name}} {{"\n"}}user hobby {{.Hobby}}{{"\n"}}`
	TemplateRun(tempText, user)
}

// {{- -}} 清除两边的空格
func TestTemplateWriteLine(t *testing.T) {
	// {{- /* a comment with white space trimmed from preceding and following text */ -}}
	user := User{Name: "lisi", Hobby: "cook"}
	tempText := `user name {{.Name}}      {{- "hello" }}     user hobby {{.Hobby}}       {{- "55"}}{{"\n"}}`
	TemplateRun(tempText, user)

}

// {{println}} 调用 println函数，{{println .}} 打印输入的上下文，也就是user对象
func TestTemplatePipeline(t *testing.T) {
	// {{pipeline}}
	user := User{Name: "lisi", Hobby: "cook"}
	tempText := `user name {{.Name}} {{println}} user hobby {{.Hobby}}  {{"\n"}}`
	TemplateRun(tempText, user)
}

// {{println}} 调用 println函数，{{println .}} 打印输入的上下文，也就是user对象
func TestTemplatePipeline2(t *testing.T) {
	// {{pipeline}}
	tempText := `{{printf "%q" .}}` // %q表示原样输出、.表示上下文，也就是入参
	TemplateRun(tempText, "current value \n is andy")
}
// {{if .xxx}} {{else}} {{end}} 这是表达式里面if语句
func TestTemplatePipelineEnd(t *testing.T) {
	// {{if pipeline}} T1 {{end}}
	user := User{Name: "lisi", Hobby: "cook"}
	tempText := `{{if .Age}}  age is {{.Age}} {{ else }} default is 18  {{end}}`
	TemplateRun(tempText, user)
}

// {{range .}} {{end}} 这里表达的是循环，其中的.(dot)表示当前上下文的数据
func TestTemplateRange(t *testing.T) {
	// {{range pipeline}} T1 {{end}}
	nameList := []string{"Andy", "bob", "Alise"}
	// range后面的.表示入参，里面的.表示的是遍历后的item
	tempText := `{{range .}} {{"\n"}} - {{.}} {{end}}{{"\n"}}`
	TemplateRun(tempText, nameList)
}

func TemplateRun(tempText string, data any) {
	parse, err := template.New("temp").Parse(tempText)
	if err != nil {
		log.Fatal(err)
	}
	err = parse.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
