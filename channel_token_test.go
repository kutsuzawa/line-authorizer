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
	cases := []struct {
		name      string
		inputFile string
		expect    string
	}{
		{name: "Successful publishing channel token", inputFile: "testdata/channel_auth_token.json", expect: "W1TeHCgfH2Liwa"},
		{name: "Failed publishing channel token", inputFile: "testdata/channel_auth_token.json", expect: "W1TeHCgfH2Liwa"},
	}

	for _, c := range cases {
		ts := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/v2/oauth/accessToken" {
					w.WriteHeader(http.StatusNotFound)
					return
				}
				b, _ := ioutil.ReadFile(c.inputFile)
				w.Write(b)
				return
			},
		))

		t.Run(c.name, func(t *testing.T) {
			config := authorizer.Config{
				APIAddress: ts.URL,
			}
			client := authorizer.NewClient(config)
			token, err := client.PublishChannelToken()
			if err != nil {
				t.Errorf("err shoud not occur, but err: %v", err)
			}
			if *token != c.expect {
				t.Errorf("token should be %s, but actual is %s", c.expect, *token)
			}
		})
		ts.Close()
	}
}
