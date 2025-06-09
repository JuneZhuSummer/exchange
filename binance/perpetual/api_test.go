package perpetual

import "testing"

var (
	endpoint = TestUrl
	key      = ""
	secret   = ""
)

func TestNew(t *testing.T) {
	i := New(endpoint, key, secret)

	t.Logf("%+v", i)
}
