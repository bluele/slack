package slack

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"time"
)

// API channels.list: Lists all channels in a Slack team.
func (sl *Slack) ChannelsList() ([]*Channel, error) {
	uv := sl.urlValues()
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

// response type for `channels.list` api
type ChannelsListAPIResponse struct {
	BaseAPIResponse
	RawChannels json.RawMessage `json:"channels"`
}

// slack channel type
type Channel struct {
	Id          string          `json:"id"`
	Name        string          `json:"name"`
	IsChannel   bool            `json:"is_channel"`
	Created     int             `json:"created"`
	Creator     string          `json:"creator"`
	IsArchived  bool            `json:"is_archived"`
	IsGeneral   bool            `json:"is_general"`
	IsMember    bool            `json:"is_member"`
	Members     []string        `json:"members"`
	RawTopic    json.RawMessage `json:"topic"`
	RawPurpose  json.RawMessage `json:"purpose"`
	NumMembers  int             `json:"num_members"`
	LastRead    string          `json:"last_read,omitempty"`
	UnreadCount float64         `json:"unread_count,omitempty"`
}

// Channels returns a slice of channel object from a response of `channels.list` api.
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

// response type for `channels.list` api
type ChannelsInfoAPIResponse struct {
	BaseAPIResponse
	Channel Channel
}

func (sl *Slack) ChannelsInfo(id string) (*Channel, error) {
	uv := sl.urlValues()
	uv.Add("channel", id)
	body, err := sl.GetRequest(channelsInfoApiEndpoint, uv)
	if err != nil {
		return nil, err
	}
	res := new(ChannelsInfoAPIResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, errors.New(res.Error)
	}
	return &res.Channel, nil
}

// FindChannel returns a channel object that satisfy conditions specified.
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

// FindChannelByName returns a channel object that matches name specified.
func (sl *Slack) FindChannelByName(name string) (*Channel, error) {
	return sl.FindChannel(func(channel *Channel) bool {
		return channel.Name == name
	})
}

// ChannelsMark moves the read cursor for the chosen channel.
func (sl *Slack) ChannelsMark(name, ts string) error {
	uv := sl.urlValues()
	uv.Add("name", name)
	uv.Add("ts", ts)

	_, err := sl.GetRequest(channelsMarkApiEndpoint, uv)
	if err != nil {
		return err
	}
	return nil
}

// API channels.join: Joins a channel, creating it if needed.
func (sl *Slack) JoinChannel(name string) error {
	uv := sl.urlValues()
	uv.Add("name", name)

	_, err := sl.GetRequest(channelsJoinApiEndpoint, uv)
	if err != nil {
		return err
	}
	return nil
}

type Message struct {
	Type    string `json:"type"`
	Subtype string `json:"subtype"`
	Ts      string `json:"ts"`
	UserId  string `json:"user"`
	Text    string `json:"text"`
}

func (msg *Message) Timestamp() *time.Time {
	tsf, _ := strconv.ParseFloat(msg.Ts, 64)
	ts := time.Unix(int64(tsf), 0)
	return &ts
}

// option type for `channels.history` api
type ChannelsHistoryOpt struct {
	Channel   string  `json:"channel"`
	Latest    float64 `json:"latest"`
	Oldest    float64 `json:"oldest"`
	Inclusive int     `json:"inclusive"`
	Count     int     `json:"count"`
}

func (opt *ChannelsHistoryOpt) Bind(uv *url.Values) error {
	uv.Add("channel", opt.Channel)
	if opt.Latest != 0.0 {
		uv.Add("lastest", strconv.FormatFloat(opt.Latest, 'f', 6, 64))
	}
	if opt.Oldest != 0.0 {
		uv.Add("oldest", strconv.FormatFloat(opt.Oldest, 'f', 6, 64))
	}
	uv.Add("inclusive", strconv.Itoa(opt.Inclusive))
	if opt.Count != 0 {
		uv.Add("count", strconv.Itoa(opt.Count))
	}
	return nil
}

// response type for `channels.history` api
type ChannelsHistoryResponse struct {
	BaseAPIResponse
	Latest   float64    `json:"latest"`
	Messages []*Message `json:"messages"`
	HasMore  bool       `json:"has_more"`
}

// API channels.history: Fetches history of messages and events from a channel.
func (sl *Slack) ChannelsHistory(opt *ChannelsHistoryOpt) ([]*Message, error) {
	uv := sl.urlValues()
	err := opt.Bind(uv)
	if err != nil {
		return nil, err
	}
	body, err := sl.GetRequest(channelsHistoryApiEndpoint, uv)
	if err != nil {
		return nil, err
	}
	res := new(ChannelsHistoryResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, errors.New(res.Error)
	}
	return res.Messages, nil
}
