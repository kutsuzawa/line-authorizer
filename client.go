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
	HTTPClient *http.Client
	APIAddress string
}

// DefaultConfig configuration for client
func DefaultConfig() *Config {
	return &Config{
		ID:         "admin",
		Secret:     "admin",
		HTTPClient: http.DefaultClient,
		APIAddress: "https://api.line.me",
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
	if config.HTTPClient == nil {
		config.HTTPClient = defConfig.HTTPClient
	}
	if len(config.APIAddress) == 0 {
		config.APIAddress = defConfig.APIAddress
	}
	return &Client{
		config,
	}
}

// NewRequest is used to create a new Request
func (c *Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	u := fmt.Sprintf("%s/%s", c.APIAddress, path)
	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Do sends an HTTP request
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.HTTPClient.Do(req)
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
