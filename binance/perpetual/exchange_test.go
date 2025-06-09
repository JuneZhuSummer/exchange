package perpetual

import (
	"context"
	"exchange/binance"
	"testing"
)

func TestExchangeInfo(t *testing.T) {
	i := New(endpoint, key, secret)

	ctx := context.Background()
	in := &binance.Empty{}
	reply, err := i.ExchangeInfo(ctx, in)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", reply)
}
