package slack

import (
	"encoding/json"
	"errors"
	"net/url"
)

// API chat.postMessage: Sends a message to a channel.
func (sl *Slack) ChatPostMessage(channel string, text string, opt *ChatPostMessageOpt) error {
	uv := sl.buildChatPostMessageUrlValues(opt)
	uv.Add("channel", channel)
	uv.Add("text", text)

	body, err := sl.GetRequest(chatPostMessageApiEndpoint, uv)
	if err != nil {
		return err
	}
	res := new(ChatPostMessageAPIResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return err
	}
	if !res.Ok {
		return errors.New(res.Error)
	}
	return nil
}

type ChatPostMessageOpt struct {
	Parse       string
	LinkNames   string
	AttachMents string
	UnfurlLinks string
	UnfurlMedia string
	IconUrl     string
	IconEmoji   string
}

type ChatPostMessageAPIResponse struct {
	BaseAPIResponse
	Channel string `json:"channel"`
	Ts      string `json:"ts"`
}

func (sl *Slack) buildChatPostMessageUrlValues(opt *ChatPostMessageOpt) *url.Values {
	uv := sl.UrlValues()
	if opt == nil {
		return uv
	}

	if opt.Parse != "" {
		uv.Add("parse", opt.Parse)
	}
	if opt.LinkNames != "" {
		uv.Add("link_names", opt.LinkNames)
	}
	if opt.AttachMents != "" {
		uv.Add("attachments", opt.AttachMents)
	}
	if opt.UnfurlLinks != "" {
		uv.Add("unfurl_links", opt.UnfurlLinks)
	}
	if opt.UnfurlMedia != "" {
		uv.Add("unfurl_media", opt.UnfurlMedia)
	}
	if opt.IconUrl != "" {
		uv.Add("icon_url", opt.IconUrl)
	}
	if opt.IconEmoji != "" {
		uv.Add("icon_emoji", opt.IconEmoji)
	}

	return uv
}
