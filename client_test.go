package authorizer_test

import (
	"net/http"
	"testing"

	"github.com/kutsuzawa/line-authorizer"
)

func TestDefaultConfig(t *testing.T) {
	t.Helper()
	t.Run("Default Config", func(t *testing.T) {
		c := authorizer.DefaultConfig()
		if c.ID != "admin" {
			t.Errorf("ID should be %s, but actual is %s", "admin", c.ID)
		}
		if c.Secret != "admin" {
			t.Errorf("Secret should be %s, but actual is %s", "admin", c.Secret)
		}
		if c.HTTPClient != http.DefaultClient {
			t.Errorf("HTTPClient should be default, but actual is %v", c.HTTPClient)
		}
	})
}
