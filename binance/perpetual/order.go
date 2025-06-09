package perpetual

import (
	"context"

	jsoniter "github.com/json-iterator/go"
)

// Order 下单
func (i *API) Order(ctx context.Context, req *OrderRequest) (reply *OrderReply, err error) {
	url := i.endpoint + "/fapi/v1/order"

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

// BatchOrder 批量下单
func (i *API) BatchOrder(ctx context.Context, req *BatchOrderRequest) (reply *BatchOrderReply, err error) {
	url := i.endpoint + "/fapi/v1/batchOrders"

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

// CancelOrder 撤单
func (i *API) CancelOrder(ctx context.Context, req *CancelOrderRequest) (reply *CancelOrderReply, err error) {
	url := i.endpoint + "/fapi/v1/order"

	bts, err := i.delete(ctx, url, req)
	if err != nil {
		return nil, err
	}
	err = jsoniter.Unmarshal(bts, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// BatchCancelOrder 批量撤单
func (i *API) BatchCancelOrder(ctx context.Context, req *BatchCancelOrderRequest) (reply *BatchCancelOrderReply, err error) {
	url := i.endpoint + "/fapi/v1/batchOrders"

	bts, err := i.delete(ctx, url, req)
	if err != nil {
		return nil, err
	}
	err = jsoniter.Unmarshal(bts, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// BatchCancelAllOrder 批量全部撤单
func (i *API) BatchCancelAllOrder(ctx context.Context, req *BatchCancelAllOrderRequest) (reply *BatchCancelAllOrderReply, err error) {
	url := i.endpoint + "/fapi/v1/allOpenOrders"

	bts, err := i.delete(ctx, url, req)
	if err != nil {
		return nil, err
	}
	err = jsoniter.Unmarshal(bts, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// FetchOneOrder 查询单个订单
func (i *API) FetchOneOrder(ctx context.Context, req *FetchOneOrderRequest) (reply *FetchOneOrderReply, err error) {
	url := i.endpoint + "/fapi/v1/order"

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

// FetchOpenOrder 查询当前所有挂单
func (i *API) FetchOpenOrder(ctx context.Context, req *FetchOpenOrderRequest) (reply *FetchOpenOrderReply, err error) {
	url := i.endpoint + "/fapi/v1/openOrders"

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
