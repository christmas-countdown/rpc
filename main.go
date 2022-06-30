package main

import (
	"fmt"
	"math"
	"time"

	"github.com/hugolgst/rich-go/client"
)

func setActivity() {
	format := "2006-01-02T15:04:05"
	christmas, err := time.Parse(format, fmt.Sprintf("%d-12-25T00:00:00", time.Now().Year()))
	if err != nil {
		panic(err)
	}
	now := time.Now()
	diff := christmas.Unix() - now.Unix()
	sleeps := int64(math.Ceil(float64(diff) / 60 / 60 / 24))
	err = client.SetActivity(client.Activity{
		State:      "until Christmas",
		Details:    fmt.Sprintf("%d sleeps left", sleeps),
		LargeImage: "santa",
		Buttons: []*client.Button{
			{
				Label: "Watch the live countdown",
				Url:   "https://christmascountdown.live",
			},
			{
				Label: "Add the bot",
				Url:   "https://discord.com/oauth2/authorize?client_id=509851616216875019&permissions=537149440&scope=applications.commands%20bot",
			},
		},
	})
	t := now.Format(time.RFC3339)
	fmt.Printf("[%s] Updated activity (%d sleeps left)\n", t, sleeps)
}

func main() {
	err := client.Login("509851616216875019")
	if err != nil {
		panic(err)
	}
	setActivity()
	for range time.Tick(time.Hour) {
		go setActivity()
	}
}
