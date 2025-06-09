package perpetual

import (
	"context"
	"exchange/binance"

	jsoniter "github.com/json-iterator/go"
)

func (i *API) ExchangeInfo(ctx context.Context, req *binance.Empty) (reply *ExchangeInfoReply, err error) {
	url := i.endpoint + "/fapi/v1/exchangeInfo"

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
