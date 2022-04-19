package test

import (
	"fmt"
	"testing"
	"time"
)
import "github.com/robfig/cron/v3"

func TestCron(t *testing.T) {
	c := cron.New(cron.WithSeconds())
	// c := cron.New()

	c.AddFunc("* * * * * *", func() {
		fmt.Println(time.Now())
	})

	c.Start()

	select {}
}
