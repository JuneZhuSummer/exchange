package perpetual

import "testing"

var (
	endpoint = "https://testnet.binancefuture.com"
	key      = "378803691eacbb6952f1aa0ea96e02a19a2dd11ab646758b265ae72feffbb79a"
	secret   = "7fb148bc2c73f8da5ef61516287517119425bc524fed08068fe19fd3eddc7b1e"
)

func TestNew(t *testing.T) {
	i := New(endpoint, key, secret)

	t.Logf("%+v", i)
}
