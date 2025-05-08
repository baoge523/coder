package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	cron := cron.New()


	// cronstr := "0 */1 * * * *"  // "0 0/1 * * * *"

	teststr := "0 0/5 * * * *"  // "0 */5 * * *"
	cron.AddFunc(teststr, func() {
		fmt.Printf("time = %s,hello world \n",time.Now().String())
	})

	cron.Start()
	time.Sleep(10*time.Minute)
	fmt.Println("success")
}
