package slack

import (
	"encoding/json"
	"errors"
)

type Member struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Deleted  bool   `json:"deleted"`
	Color    string `json:"color"`
	Profile  *ProfileInfo
	IsAdmin  bool `json:"is_admin"`
	IsOwner  bool `json:"is_owner"`
	Has2fa   bool `json:"has_2fa"`
	HasFiles bool `json:"has_files"`
}

type ProfileInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	RealName  string `json:"real_name"`
	Email     string `json:"email"`
	Skype     string `json:"skype"`
	Phone     string `json:"phone"`
	Image24   string `json:"image_24"`
	Image32   string `json:"image_32"`
	Image48   string `json:"image_48"`
	Image72   string `json:"image_72"`
	Image192  string `json:"image_192"`
}

func (sl *Slack) UsersList() ([]*Member, error) {
	uv := sl.UrlValues()
	body, err := sl.GetRequest(usersListApiEndpoint, uv)
	if err != nil {
		return nil, err
	}
	res := new(UsersListAPIResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, errors.New(res.Error)
	}
	return res.Members()
}

type UsersListAPIResponse struct {
	BaseAPIResponse
	RawMembers json.RawMessage `json:"members"`
}

func (res *UsersListAPIResponse) Members() ([]*Member, error) {
	var members []*Member
	err := json.Unmarshal(res.RawMembers, &members)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (sl *Slack) FindUser(cb func(*Member) bool) (*Member, error) {
	members, err := sl.UsersList()
	if err != nil {
		return nil, err
	}
	for _, member := range members {
		if cb(member) {
			return member, nil
		}
	}
	return nil, errors.New("No such user.")
}
