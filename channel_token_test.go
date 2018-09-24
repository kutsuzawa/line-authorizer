package authorizer_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kutsuzawa/line-authorizer"
)

func TestChannelToken_Publish(t *testing.T) {
	t.Helper()
	t.Run("Successful publishing channel token", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/v2/oauth/accessToken" {
					w.WriteHeader(http.StatusNotFound)
					return
				}
				b, _ := ioutil.ReadFile("testdata/channel_auth_token.json")
				w.Write(b)
				return
			},
		))
		defer ts.Close()

		config := authorizer.Config{
			APIAddress: ts.URL,
		}
		client := authorizer.NewClient(config)
		token, err := client.PublishChannelToken()
		if err != nil {
			t.Errorf("err shoud not occur, but err: %v", err)
		}
		if *token != "W1TeHCgfH2Liwa" {
			t.Errorf("token should be %s, but actual is %s", "W1TeHCgfH2Liwa", *token)
		}
	})

	t.Run("Failed publishing channel token", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/v2/oauth/accessToken" {
					w.WriteHeader(http.StatusNotFound)
					return
				}
				b, _ := ioutil.ReadFile("testdata/channel_auth_token_err.json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write(b)
				return
			},
		))
		defer ts.Close()

		config := authorizer.Config{
			APIAddress: ts.URL,
		}
		client := authorizer.NewClient(config)
		token, err := client.PublishChannelToken()
		if token != nil {
			t.Errorf("token should be nil, but actual is %v", token)
		}
		if err == nil {
			t.Errorf("err shoud occur, but it does not")
		}
	})
}
