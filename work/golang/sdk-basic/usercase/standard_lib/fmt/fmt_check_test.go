package fmt_all

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	sqlFmtApplicationPolicy := "SELECT uniqueId,appAddress,groupId," +
		"information,isShielded FROM rApplicationPolicy WHERE isShielded = 0 AND groupId %% %d = %d"
	sprintf := fmt.Sprintf(sqlFmtApplicationPolicy, 1, 2)
	fmt.Println(sprintf)
}
