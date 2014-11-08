package slack

import (
	"encoding/json"
	"errors"
)

const channelsListApiEndpoint = "channels.list"

// API channels.list: Lists all channels in a Slack team.
func (sl *Slack) ChannelsList() ([]Channel, error) {
	uv := sl.UrlValues()
	body, err := sl.GetRequest(channelsListApiEndpoint, uv)
	if err != nil {
		return nil, err
	}
	res := new(ChannelsListAPIResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, errors.New(res.Error)
	}
	return res.Channels()
}

type ChannelsListAPIResponse struct {
	Ok          bool            `json:"ok"`
	RawChannels json.RawMessage `json:"channels"`
	Error       string          `json:"error"`
}

type Channel struct {
	Id         string          `json:"id"`
	Name       string          `json:"name"`
	IsChannel  bool            `json:"is_channel"`
	Created    int             `json:"created"`
	Creator    string          `json:"creator"`
	IsArchived bool            `json:"is_archived"`
	IsGeneral  bool            `json:"is_general"`
	IsMember   bool            `json:"is_member"`
	Members    []string        `json:"members"`
	RawTopic   json.RawMessage `json:"topic"`
	RawPurpose json.RawMessage `json:"purpose"`
	NumMembers int             `json:"num_members"`
}

type Topic struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet string `json:"last_set"`
}

type Purpose struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet string `json:"last_set"`
}

func (res *ChannelsListAPIResponse) Channels() ([]Channel, error) {
	var chs []Channel
	err := json.Unmarshal(res.RawChannels, &chs)
	if err != nil {
		return nil, err
	}
	return chs, nil
}

func (ch *Channel) Topic() (*Topic, error) {
	tp := new(Topic)
	err := json.Unmarshal(ch.RawTopic, tp)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func (ch *Channel) Purpose() (*Purpose, error) {
	pp := new(Purpose)
	err := json.Unmarshal(ch.RawPurpose, pp)
	if err != nil {
		return nil, err
	}
	return pp, nil
}
