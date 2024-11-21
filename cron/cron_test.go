package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"testing"
)

func Test_cron(t *testing.T) {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("@every 10s", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()

	select {}

	c.Stop()
}
