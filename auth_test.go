package auth_test

import (
	"testing"

	"github.com/kutsuzawa/line-authorizer"
)

func TestDummy(t *testing.T) {
	t.Helper()
	t.Run("dummy", func(t *testing.T) {
		expect := "foo"
		dummy := auth.Dummy()
		if dummy != expect {
			t.Errorf("dummy string shoud be %s, but actual is %s", expect, dummy)
		}
	})
}
