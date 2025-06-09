package perpetual

import (
	"context"
	"testing"
	"time"
)

func TestPosition(t *testing.T) {
	i := New(endpoint, key, secret)

	ctx := context.Background()
	in := &PositionRequest{
		Symbol:     "ETHUSDT",
		RecvWindow: 5000,
		Timestamp:  time.Now().UnixMilli(),
	}
	reply, err := i.Position(ctx, in)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", reply)
}

func TestPositionSideDual(t *testing.T) {
	i := New(endpoint, key, secret)
	ctx := context.Background()
	in := &PositionSideDualRequest{
		DualSidePosition: true,
		RecvWindow:       5000,
		Timestamp:        time.Now().UnixMilli(),
	}
	reply, err := i.PositionSideDual(ctx, in)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", reply)
}
