package perpetual

import (
	"context"

	jsoniter "github.com/json-iterator/go"
)

func (i *API) Balance(ctx context.Context, req *BalanceRequest) (reply *BalanceReply, err error) {
	url := i.endpoint + "/fapi/v3/balance"

	bts, err := i.get(ctx, url, req)
	if err != nil {
		return nil, err
	}
	err = jsoniter.Unmarshal(bts, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (i *API) Account(ctx context.Context, req *AccountRequest) (reply *AccountReply, err error) {
	url := i.endpoint + "/fapi/v2/account"

	bts, err := i.get(ctx, url, req)
	if err != nil {
		return nil, err
	}
	err = jsoniter.Unmarshal(bts, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
