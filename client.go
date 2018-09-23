package authorizer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client is used to get info from line api
type Client struct {
	Config
}

// Config is used to configure the creation of a client
type Config struct {
	ID         string
	Secret     string
	HttpClient *http.Client
	ApiAddress string
}

func DefaultConfig() *Config {
	return &Config{
		ID:         "admin",
		Secret:     "admin",
		HttpClient: http.DefaultClient,
		ApiAddress: "https://api.line.me",
	}
}

// NewClient returns a new client
func NewClient(config Config) *Client {
	defConfig := DefaultConfig()
	if len(config.ID) == 0 {
		config.ID = defConfig.ID
	}
	if len(config.Secret) == 0 {
		config.Secret = defConfig.Secret
	}
	if config.HttpClient == nil {
		config.HttpClient = defConfig.HttpClient
	}
	if len(config.ApiAddress) == 0 {
		config.ApiAddress = defConfig.ApiAddress
	}
	return &Client{
		config,
	}
}

// NewRequest is used to create a new Request
func (c *Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	u := fmt.Sprintf("%s/%s", c.ApiAddress, path)
	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) decodeBody(obj interface{}, body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&obj); err != nil {
		return err
	}
	return nil
}
