package template_handle

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"text/template"
)

func TestJson(t *testing.T) {

	temp := `{"policyId": "{{.policyId}}}"`

	_, err := template.New("test").Parse(temp)
	if err != nil {
		fmt.Println("parse err")
	}

	fmt.Println("ok")
}

/*
占位符是一个json对象，期望渲染后，变成一个json
*/
func TestJson22(t *testing.T) {

	// content := `{"alarmObj": {{.dim}}}`  // {"alarmObj": {\"aaa\":\"aaa\",\"bbb\":\"bbb\"}}
	content := `{"alarmObj": "{{.dim}}"}` // {"alarmObj": "{\"aaa\":\"aaa\",\"bbb\":\"bbb\"}"}

	temp, err := template.New("test").Parse(content)
	if err != nil {
		fmt.Println("parse err")
	}
	map1 := make(map[string]interface{})

	m := make(map[string]interface{})
	m["aaa"] = "aaa"
	m["bbb"] = "bbb"

	bytes, _ := json.Marshal(m)

	replaceAll := strings.ReplaceAll(string(bytes), "\"", "\\\"")

	map1["dim"] = replaceAll

	temp.Execute(os.Stdout, map1)
}

const a = `{"content":"  {\n \"dimensions\": {\"objId\":\"f39edcc0-8d80-47c2-8f15-3054db1b46f8\",\"objName\":\"10.0.0.8#67210\",\"unInstanceId\":\"ins-6zvvt4fb\"},\n   \"uin\": \"110000000578\",\n    \"appId\": \"1255000457\",\n    \"sessionId\": \"Ezd1zGTJwRAEZxvDodEin15B\",\n    \"alarmStatus\": 1,\n     \"region\": \"ap-shenzhen-hqtest-ops\",\n     \"namespace\": \"qce/cvm\",\n       \"policyId\": \"policy-nih4czn7\",\n     \"policyType\": \"cvm_device\",\n     \"policyName\": \"测试123\",\n      \"metricName\": \"cpu_usage\",\n      \"metricShowName\": \"CPU利用率 \",\n      \"calcType\": \">\",\n      \"calcValue\": \"0\",\n      \"currentValue\": \"1.483\",\n      \"unit\": \"%\",\n       \"period\": \"60\",\n       \"periodNum\": \"1\",\n       \"alarmNotifyPeriod\": 1,\n       \"firstOccurTime\": \"2025-02-21 12:33:00\",\n       \"durationTime\": 3000,\n       \"recoverTime\": \"0\"\n  }"}`

//  \"dimensions\": {\"objId\":\"f39edcc0-8d80-47c2-8f15-3054db1b46f8\",\"objName\":\"10.0.0.8#67210\",\"unInstanceId\":\"ins-6zvvt4fb\"},\n

const b = `{"content":"{\n \"policyId\": \"aaa\",\n \"region\": \"111\"\n}"}`

func TestJson3(t *testing.T) {
	replaceAll := strings.ReplaceAll(a, "\\n", "")
	fmt.Println(replaceAll)
	m1 := make(map[string]string)
	err := json.Unmarshal([]byte(replaceAll), &m1)
	if err != nil {
		fmt.Println(err)
	}

}

func TestJson4(t *testing.T) {

	s := "1255000036"
	atoi, _ := strconv.Atoi(s)
	fmt.Println(atoi)

}
