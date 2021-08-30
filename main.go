package main

import (
	"fmt"
	"github.com/nshipman-io/pr-discord-bot/bot"
	"github.com/nshipman-io/pr-discord-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Starting Discord bot...")
	bot.Start()
	<-make(chan struct{})
	return
}
