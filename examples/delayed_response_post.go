package main

import (
	"github.com/bluele/slack"
)

// Please change these values to suit your environment
const (
	responseURL = "https://hooks.slack.com/services/xxxxxx/xxxxxx/xxxxxxxxxxxxxx"
)

func main() {
	response := slack.NewDelayedResponse(responseURL)
	err := response.PostDelayedResponse(&slack.ResponsePayload{
		Text: "hello!",
    ResponseType: "in_channel",
		Attachments: []*slack.Attachment{
			{Text: "danger", Color: "danger"},
		},
	})
	if err != nil {
		panic(err)
	}
}

