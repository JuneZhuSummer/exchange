package perpetual

import (
	"context"
	"testing"
	"time"
)

func TestBalance(t *testing.T) {
	i := New(endpoint, key, secret)

	ctx := context.Background()
	in := &BalanceRequest{
		RecvWindow: 5000,
		Timestamp:  time.Now().UnixMilli(),
	}
	reply, err := i.Balance(ctx, in)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", reply)
}

func TestAccount(t *testing.T) {
	i := New(endpoint, key, secret)

	ctx := context.Background()
	in := &AccountRequest{
		RecvWindow: 5000,
		Timestamp:  time.Now().UnixMilli(),
	}
	reply, err := i.Account(ctx, in)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", reply)
}
