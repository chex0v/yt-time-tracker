package client

import (
	"io"
	"log"
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

func (c *Client) headerToken() string {
	return "Bearer " + c.token
}

func (c *Client) Do(method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, c.Url+path, body)

	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {c.headerToken()},
	}

	return c.HTTPClient.Do(req)
}
