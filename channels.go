package slack

import (
	"encoding/json"
	"errors"
)

// API channels.list: Lists all channels in a Slack team.
func (sl *Slack) ChannelsList() ([]*Channel, error) {
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
	BaseAPIResponse
	RawChannels json.RawMessage `json:"channels"`
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

func (res *ChannelsListAPIResponse) Channels() ([]*Channel, error) {
	var chs []*Channel
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

func (sl *Slack) FindChannel(cb func(*Channel) bool) (*Channel, error) {
	channels, err := sl.ChannelsList()
	if err != nil {
		return nil, err
	}
	for _, channel := range channels {
		if cb(channel) {
			return channel, nil
		}
	}
	return nil, errors.New("No such channel.")
}

func (sl *Slack) FindChannelByName(name string) (*Channel, error) {
	return sl.FindChannel(func(channel *Channel) bool {
		return channel.Name == name
	})
}
