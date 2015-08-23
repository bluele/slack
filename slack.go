package slack

import (
	"net/url"
)

type Slack struct {
	token string
}

func New(token string) *Slack {
	return &Slack{
		token: token,
	}
}

func (sl *Slack) UrlValues() *url.Values {
	uv := url.Values{}
	uv.Add("token", sl.token)
	return &uv
}
