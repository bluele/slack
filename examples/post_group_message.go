package main

import (
	"github.com/bluele/slack"
)

// Please change these values to suit your environment
const (
	token     = "your-api-token"
	groupName = "group-name"
)

func main() {
	api := slack.New(token)
	group, err := api.FindGroupByName(groupName)
	if err != nil {
		panic(err)
	}

	err = api.ChatPostMessage(group.Id, "Hello, world!", &slack.ChatPostMessageOpt{AsUser: true})
	if err != nil {
		panic(err)
	}
}
