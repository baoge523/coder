package debug

import (
	"runtime/debug"
	"testing"
)

func TestDebug(t *testing.T) {

	debug.SetGCPercent(60)


}
