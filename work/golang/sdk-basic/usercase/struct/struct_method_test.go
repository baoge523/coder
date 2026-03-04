package _struct

import (
	"fmt"
	"testing"
	"time"
)

type Track struct {
	Name string
	Ch   chan int
}

func (t *Track) Close() {
	if t.Ch != nil {
		fmt.Println("close channel")
		close(t.Ch)
	}
}

func (t *Track) DoSomething() {
	go func() {
		for item := range t.Ch {
			fmt.Println(item)
		}
		fmt.Println("stopped channel")
	}()
}

func Test_Name(t *testing.T) {
	newTrack := NewTrack("test")
	newTrack.DoSomething()
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		newTrack.Ch <- i
	}
	time.Sleep(1 * time.Second)
	newTrack.Close()
}

func NewTrack(name string) Track {
	return Track{
		Name: name,
		Ch:   make(chan int, 100),
	}
}
