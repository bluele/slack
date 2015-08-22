package main

import (
	"github.com/bluele/slack"
)

const (
	token       = "your-api-token"
	channelName = "general"
)

func main() {
	api := slack.New(token)
	channelID, err := api.LookupChannelID(channelName)
	if err != nil {
		panic(err)
	}
	err = api.ChatPostMessage(channelID, "Hello, world!", nil)
	if err != nil {
		panic(err)
	}
}
