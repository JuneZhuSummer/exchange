package perpetual

import (
	"context"
	"testing"
	"time"
)

func TestTrades(t *testing.T) {
	i := New(endpoint, key, secret)

	ctx := context.Background()
	in := &TradesRequest{
		Symbol:     "BTCUSDT",
		OrderId:    4065461425,
		RecvWindow: 5000,
		Timestamp:  time.Now().UnixMilli(),
	}
	reply, err := i.Trades(ctx, in)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", reply)
}
