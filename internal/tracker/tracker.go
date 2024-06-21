package tracker

import (
	"net/http"
	"time"
)

type Client struct {
	Url        string
	token      string
	HTTPClient *http.Client
}

func NewClient(url, token string) *Client {
	return &Client{
		token: token,
		Url:   url,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c Client) headerToken() string {
	return "Bearer " + c.token
}
