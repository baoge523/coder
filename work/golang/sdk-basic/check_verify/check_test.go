package check_verify

import (
	"fmt"
	"strings"
	"testing"
)

func TestCheck(t *testing.T) {
	queryPolicyByOriginIdSQL    := "SELECT `id`,`revision` FROM `e_policy` WHERE `app_id` = ? AND `tags` like `%s` " +
		"AND `deleted` = 0"
	param := fmt.Sprintf(`*"groupId":"%s"*`, "1601489")
	param = strings.ReplaceAll(param, "*", "%")
	fmt.Println(param)
	querySQL := fmt.Sprintf(queryPolicyByOriginIdSQL, param)
	fmt.Println(querySQL)
}
