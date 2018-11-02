package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type DelayedResponse struct {
	responseURL string
}

type ResponsePayload struct {
	ResponseType string        `json:"response_type,omitempty"`
	Text         string        `json:"text,omitempty"`
	UnfurlLinks  bool          `json:"unfurl_links,omitempty"`
	LinkNames    string        `json:"link_names,omitempty"`
	Attachments  []*Attachment `json:"attachments,omitempty"`
}

func NewDelayedResponse(responseURL string) *DelayedResponse {
	return &DelayedResponse{responseURL}
}

func (dr *DelayedResponse) PostDelayedResponse(payload *ResponsePayload) error {
	//Slack recommends declaring this even if you wish to use the default "ephemeral" value
	if payload.ResponseType == "" {
		payload.ResponseType = "ephemeral"
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	resp, err := http.Post(dr.responseURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(t))
	}

	return nil
}
