package perpetual

import (
	"context"
	"exchange/binance"

	jsoniter "github.com/json-iterator/go"
)

func (i *API) Position(ctx context.Context, req *PositionRequest) (reply *PositionReply, err error) {
	url := i.endpoint + "/fapi/v3/positionRisk"

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

func (i *API) PositionSideDual(ctx context.Context, req *PositionSideDualRequest) (reply *binance.Reply, err error) {
	url := i.endpoint + "/fapi/v1/positionSide/dual"
	bts, err := i.post(ctx, url, req)
	if err != nil {
		return nil, err
	}
	err = jsoniter.Unmarshal(bts, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
