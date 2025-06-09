package perpetual

import "testing"

var (
	key    = ""
	secret = ""
)

func TestNew(t *testing.T) {
	i := New(TestUrl, key, secret)

	t.Logf("%+v", i)
}
