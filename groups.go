package slack

import (
	"encoding/json"
	"errors"
)

func (sl *Slack) GroupsList() ([]*Group, error) {
	uv := sl.UrlValues()
	body, err := sl.GetRequest(groupsListApiEndpoint, uv)
	if err != nil {
		return nil, err
	}
	res := new(GroupsListAPIResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, errors.New(res.Error)
	}
	return res.Groups()
}

func (sl *Slack) CreateGroup(name string) error {
	uv := sl.UrlValues()
	uv.Add("name", name)

	_, err := sl.GetRequest(groupsCreateApiEndpoint, uv)
	if err != nil {
		return err
	}
	return nil
}

type Group struct {
	Id         string          `json:"id"`
	Name       string          `json:"name"`
	Created    int             `json:"created"`
	Creator    string          `json:"creator"`
	IsArchived bool            `json:"is_archived"`
	Members    []string        `json:"members"`
	RawTopic   json.RawMessage `json:"topic"`
	RawPurpose json.RawMessage `json:"purpose"`
}

type GroupsListAPIResponse struct {
	BaseAPIResponse
	RawGroups json.RawMessage `json:"groups"`
}

func (res *GroupsListAPIResponse) Groups() ([]*Group, error) {
	var groups []*Group
	err := json.Unmarshal(res.RawGroups, &groups)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

type GroupsCreateAPIResponse struct {
	BaseAPIResponse
	RawGroup json.RawMessage `json:"group"`
}

func (res *GroupsCreateAPIResponse) Group() (*Group, error) {
	group := Group{}
	err := json.Unmarshal(res.RawGroup, &group)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (sl *Slack) FindGroup(cb func(*Group) bool) (*Group, error) {
	groups, err := sl.GroupsList()
	if err != nil {
		return nil, err
	}
	for _, group := range groups {
		if cb(group) {
			return group, nil
		}
	}
	return nil, errors.New("No such group.")
}

func (sl *Slack) FindGroupByName(name string) (*Group, error) {
	return sl.FindGroup(func(group *Group) bool {
		return group.Name == name
	})
}

func (sl *Slack) InviteGroup(channelId, userId string) error {
	uv := sl.UrlValues()
	uv.Add("channel", channelId)
	uv.Add("user", userId)

	_, err := sl.GetRequest(channelsJoinApiEndpoint, uv)
	if err != nil {
		return err
	}
	return nil
}
