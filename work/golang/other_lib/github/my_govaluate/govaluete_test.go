package my_govaluate

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

import "testing"

func Test_1(t *testing.T) {
	expression, err := govaluate.NewEvaluableExpression("cup_usage")
	if err != nil {
		fmt.Println(err)
	}

	params := make(govaluate.MapParameters)
	params["aa"] = 1
	params["cup_usage"] = 1.11
	params["bb"] = 22

	eval, _ := expression.Eval(params)

	fmt.Printf("%+v", eval)
}
