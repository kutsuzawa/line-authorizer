package authorizer

import (
	"io"
	"net/url"
	"fmt"
)

// Client is used to get info from line api
type Client struct {
	ID     string
	Secret string
}

// Request is used to help build up a request
type Request struct {
	method string
	url    string
	params url.Values
	body   io.Reader
	obj    interface{}
}

// NewClient returns a new client
func NewClient(id, secret string) *Client {
	return &Client{
		ID:     id,
		Secret: secret,
	}
}

// NewRequest is used to create a new Request
func (c *Client) NewRequest(method, path string, body io.Reader) *Request {
	u := fmt.Sprintf("https://api.line.me/v2/oauth/accessToken/%s", path)
	r := &Request{
		method: method,
		url:    u,
		params: make(map[string][]string),
	}

	// Set request body
	r.body = body
	return r
}
