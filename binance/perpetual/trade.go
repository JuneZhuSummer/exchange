package perpetual

import (
	"context"

	jsoniter "github.com/json-iterator/go"
)

func (i *API) Trades(ctx context.Context, req *TradesRequest) (reply *TradesReply, err error) {
	url := i.endpoint + "/fapi/v1/userTrades"

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
