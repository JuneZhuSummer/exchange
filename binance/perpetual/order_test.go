package perpetual

import (
	"context"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestOrder(t *testing.T) {
	i := New(endpoint, key, secret)

	ctx := context.Background()
	in := &OrderRequest{
		Symbol:   "BTCUSDT",
		Side:     SideBuy,
		Type:     TypeMarket,
		Quantity: decimal.NewFromFloat(0.011000001),
		//Price:       decimal.NewFromFloat(50000),
		//TimeInForce: TimeInForceGTC,
		RecvWindow: 10000,
		Timestamp:  time.Now().UnixMilli(),
	}
	reply, err := i.Order(ctx, in)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", reply)
}

func TestCancelOrder(t *testing.T) {
	i := New(endpoint, key, secret)

	ctx := context.Background()
	in := &CancelOrderRequest{
		Symbol:     "BTCUSDT",
		OrderId:    4065211086,
		RecvWindow: 5000,
		Timestamp:  time.Now().UnixMilli(),
	}
	reply, err := i.CancelOrder(ctx, in)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", reply)
}

func TestFetchOneOrder(t *testing.T) {
	i := New(endpoint, key, secret)

	ctx := context.Background()
	in := &FetchOneOrderRequest{
		Symbol:     "BTCUSDT",
		OrderId:    4065212022,
		RecvWindow: 5000,
		Timestamp:  time.Now().UnixMilli(),
	}
	reply, err := i.FetchOneOrder(ctx, in)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", reply)
}
