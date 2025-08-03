package uuid

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestUuid(t *testing.T) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(newUUID.String())
}
