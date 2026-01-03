package log

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Infof("Received: %v", "Hello")
}
