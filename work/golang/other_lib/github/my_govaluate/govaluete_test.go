package my_govaluate

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"regexp"
)

import "testing"

func Test_1(t *testing.T) {
	expression, err := govaluate.NewEvaluableExpression("cup_usage")
	if err != nil {
		fmt.Println(err)
	}

	s := expression.String()
	fmt.Println(s)

	params := make(govaluate.MapParameters)
	params["aa"] = 1
	params["cup_usage"] = 1.11
	params["bb"] = 22

	eval, err := expression.Eval(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", eval)
}

var numberReplace = regexp.MustCompile(`\d+`)

func Test2(t *testing.T) {
	exprString := "7827601 || 7827602"

	exprString = numberReplace.ReplaceAllStringFunc(exprString, func(src string) string {
		return "rule_" + src
	})
	fmt.Println(exprString)
	expression, err := govaluate.NewEvaluableExpression(exprString)
	if err != nil {
		fmt.Println(err)
	}
	params := make(govaluate.MapParameters)
	params["aa"] = 1
	params["rule_7827602"] = true
	params["rule_7827601"] = false

	eval, err := expression.Eval(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", eval)
}
