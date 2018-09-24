package authorizer

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

type channelAuth struct {
	Token     string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
	TokenType string `json:"token_type"`
}

type channelAuthErr struct {
	Err            string `json:"error"`
	ErrDescription string `json:"error_description"`
}

// PublishChannelToken publishes channel access token
func (c *Client) PublishChannelToken() (*string, error) {
	body := url.Values{}
	body.Set("grant_type", "client_credentials")
	body.Set("client_id", c.ID)
	body.Set("client_secret", c.Secret)
	req, err := c.NewRequest("POST", "v2/oauth/accessToken", strings.NewReader(body.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create a HTTP Request")
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to send request: %#v", req)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusBadRequest {
		var authErr channelAuthErr
		if err := c.decodeBody(authErr, resp.Body); err != nil {
			return nil, errors.Wrap(err, "failed to decode")
		}
		return nil, errors.Errorf("return %d from line api: %s", http.StatusBadRequest, authErr.ErrDescription)
	}
	var auth channelAuth
	if err := c.decodeBody(&auth, resp.Body); err != nil {
		return nil, errors.Wrap(err, "failed to decode")
	}
	return &auth.Token, nil
}
