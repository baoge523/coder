package channel

import (
	"math/rand"
	"testing"
	"time"
)

func TestChannelSelect(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int,1)

	go func(){
		t.Log("go routine running")
		time.Sleep(time.Second * 1)
		round := rand.Intn(10)
		t.Logf("round: %d", round)
		if round < 5 {
			c1 <- 1
		} else {
			c2 <- 1
		}
	}()

	// main
	select {
    case <- c1:
		t.Log("c1 value")
    case <- c2:
		t.Log("c2 value")
    }
	t.Log("test end")
}

// select default 选项在其他case不能执行时，执行default
func TestChannelSelectDefault(t *testing.T) {
	c1 := make(chan int,1)
	c2 := make(chan int,1)

	// go routine
	go func(){
		ticker := time.NewTicker(time.Second * 1)
		select {
		case <- ticker.C:
			c2 <- 1
		}
	}()

	// main
	select {
	case <- c1:
		t.Log("c1 value")
	default:
		t.Log("default")
	}

	<- c2  // 阻塞
	t.Log("test end")
}
