package main

import (
	"fmt"
	"math"
	"time"

	"github.com/hugolgst/rich-go/client"
)

func connect() {
	err := client.Login("509851616216875019")
	if err != nil {
		fmt.Printf("[%s] %s\n", time.Now().Format(time.RFC3339), err)
		fmt.Printf("[%s] Connection failed, retrying in 30 seconds\n", time.Now().Format(time.RFC3339))
		time.Sleep(30 * time.Second)
		connect()
	} else {
		setActivity()
		for range time.Tick(time.Hour) {
			go setActivity()
		}
	}
}

func setActivity() {
	now := time.Now()
	t := now.Format(time.RFC3339)
	format := "2006-01-02T15:04:05"
	year := now.Year()
	if now.Month() == 12 && now.Day() > 25 {
		year += 1
	}
	christmas, err := time.Parse(format, fmt.Sprintf("%d-12-25T00:00:00", year))
	if err != nil {
		panic(err)
	}
	diff := christmas.Unix() - now.Unix()
	sleeps := int64(math.Ceil(float64(diff) / 60 / 60 / 24))
	if sleeps > 1 {
	        details := fmt.Sprintf("%d sleeps left", sleeps)
        } else {
		details := fmt.Sprintf("%d sleep left", sleeps)
	}
	state := "until Christmas"
	if now.Month() == 12 && now.Day() == 25 {
		details = "It's Christmas Day!"
		state = "Merry Christmas!"
	}
	err = client.SetActivity(client.Activity{
		State:      state,
		Details:    details,
		LargeImage: "santa",
		Buttons: []*client.Button{
			{
				Label: "Watch the live countdown",
				Url:   "https://christmascountdown.live/?utm_source=rpc&utm_medium=button&utm_campaign=website",
			},
			{
				Label: "Add the bot",
				Url:   "https://christmascountdown.live/discord?utm_source=rpc&utm_medium=button&utm_campaign=bot",
			},
		},
	})
	if err != nil {
		fmt.Printf("[%s] %s\n", time.Now().Format(time.RFC3339), err)
		fmt.Printf("[%s] Failed to update activity\n", t)
	} else {
		fmt.Printf("[%s] Updated activity (%d sleeps left)\n", t, sleeps)
	}
}

func main() {
	connect()
}
