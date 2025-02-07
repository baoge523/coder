package sync

import (
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {

	var m sync.Mutex

	m.Lock()
	defer m.Unlock()
}
