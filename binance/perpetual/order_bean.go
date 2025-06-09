package perpetual

import "github.com/shopspring/decimal"

const (
	SideBuy  = "BUY"
	SideSell = "SELL"
)

const (
	TypeLimit              = "LIMIT"
	TypeMarket             = "MARKET"
	TypeStop               = "STOP"
	TypeTakeProfit         = "TAKE_PROFIT"
	TypeStopMarket         = "STOP_MARKET"
	TypeTakeProfitMarket   = "TAKE_PROFIT_MARKET"
	TypeTrailingStopMarket = "TRAILING_STOP_MARKET"
)

const (
	TimeInForceGTC = "GTC"
	TimeInForceGTD = "GTD"
	TimeInForceGTX = "GTX"
	TimeInForceIOC = "IOC"
	TimeInForceFOK = "FOK"
)

const (
	StatusNew             = "NEW"
	StatusPartiallyFilled = "PARTIALLY_FILLED"
	StatusFilled          = "FILLED"
	StatusCanceled        = "CANCELED"
	StatusRejected        = "REJECTED"
	StatusExpired         = "EXPIRED"
	StatusExpiredInMatch  = "EXPIRED_IN_MATCH"
)

const (
	PositionSideBoth  = "BOTH"
	PositionSideLong  = "LONG"
	PositionSideShort = "SHORT"
)

type OrderRequest struct {
	Symbol                  string          `json:"symbol"`
	Side                    string          `json:"side"`
	PositionSide            string          `json:"positionSide"`
	Type                    string          `json:"type"`
	ReduceOnly              string          `json:"reduceOnly"`
	Quantity                decimal.Decimal `json:"quantity"`
	Price                   decimal.Decimal `json:"price"`
	NewClientOrderId        string          `json:"newClientOrderId"`
	StopPrice               decimal.Decimal `json:"stopPrice"`
	ClosePosition           string          `json:"closePosition"`
	ActivationPrice         decimal.Decimal `json:"activationPrice"`
	CallbackRate            decimal.Decimal `json:"callbackRate"`
	TimeInForce             string          `json:"timeInForce"`
	WorkingType             string          `json:"workingType"`
	PriceProtect            string          `json:"priceProtect"`
	NewOrderRespType        string          `json:"newOrderRespType"`
	PriceMatch              string          `json:"priceMatch"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
	GoodTillDate            int64           `json:"goodTillDate"`
	RecvWindow              int64           `json:"recvWindow"`
	Timestamp               int64           `json:"timestamp"`
}

type OrderReply struct {
	ClientOrderId           string `json:"clientOrderId"`
	CumQty                  string `json:"cumQty"`
	CumQuote                string `json:"cumQuote"`
	ExecutedQty             string `json:"executedQty"`
	OrderId                 int    `json:"orderId"`
	AvgPrice                string `json:"avgPrice"`
	OrigQty                 string `json:"origQty"`
	Price                   string `json:"price"`
	ReduceOnly              bool   `json:"reduceOnly"`
	Side                    string `json:"side"`
	PositionSide            string `json:"positionSide"`
	Status                  string `json:"status"`
	StopPrice               string `json:"stopPrice"`
	ClosePosition           bool   `json:"closePosition"`
	Symbol                  string `json:"symbol"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	OrigType                string `json:"origType"`
	ActivatePrice           string `json:"activatePrice"`
	PriceRate               string `json:"priceRate"`
	UpdateTime              int64  `json:"updateTime"`
	WorkingType             string `json:"workingType"`
	PriceProtect            bool   `json:"priceProtect"`
	PriceMatch              string `json:"priceMatch"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	GoodTillDate            int64  `json:"goodTillDate"`
}

type BatchOrderRequestMeta struct {
	Symbol                  string          `json:"symbol"`
	Side                    string          `json:"side"`
	PositionSide            string          `json:"positionSide"`
	Type                    string          `json:"type"`
	ReduceOnly              string          `json:"reduceOnly"`
	Quantity                decimal.Decimal `json:"quantity"`
	Price                   decimal.Decimal `json:"price"`
	NewClientOrderId        string          `json:"newClientOrderId"`
	StopPrice               decimal.Decimal `json:"stopPrice"`
	ClosePosition           string          `json:"closePosition"`
	ActivationPrice         decimal.Decimal `json:"activationPrice"`
	CallbackRate            decimal.Decimal `json:"callbackRate"`
	TimeInForce             string          `json:"timeInForce"`
	WorkingType             string          `json:"workingType"`
	PriceProtect            string          `json:"priceProtect"`
	NewOrderRespType        string          `json:"newOrderRespType"`
	PriceMatch              string          `json:"priceMatch"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
	GoodTillDate            int64           `json:"goodTillDate"`
}

type BatchOrderRequest struct {
	BatchOrders []BatchOrderRequestMeta `json:"batchOrders"`
	RecvWindow  int64                   `json:"recvWindow"`
	Timestamp   int64                   `json:"timestamp"`
}

type BatchOrderReply []OrderReply

type CancelOrderRequest struct {
	Symbol            string `json:"symbol"`
	OrderId           int64  `json:"orderId"`
	OrigClientOrderId string `json:"origClientOrderId"`
	RecvWindow        int64  `json:"recvWindow"`
	Timestamp         int64  `json:"timestamp"`
}

type CancelOrderReply struct {
	ClientOrderId           string `json:"clientOrderId"`
	CumQty                  string `json:"cumQty"`
	CumQuote                string `json:"cumQuote"`
	ExecutedQty             string `json:"executedQty"`
	OrderId                 int    `json:"orderId"`
	OrigQty                 string `json:"origQty"`
	Price                   string `json:"price"`
	ReduceOnly              bool   `json:"reduceOnly"`
	Side                    string `json:"side"`
	PositionSide            string `json:"positionSide"`
	Status                  string `json:"status"`
	StopPrice               string `json:"stopPrice"`
	ClosePosition           bool   `json:"closePosition"`
	Symbol                  string `json:"symbol"`
	TimeInForce             string `json:"timeInForce"`
	OrigType                string `json:"origType"`
	Type                    string `json:"type"`
	ActivatePrice           string `json:"activatePrice"`
	PriceRate               string `json:"priceRate"`
	UpdateTime              int64  `json:"updateTime"`
	WorkingType             string `json:"workingType"`
	PriceProtect            bool   `json:"priceProtect"`
	PriceMatch              string `json:"priceMatch"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	GoodTillDate            int    `json:"goodTillDate"`
}

type BatchCancelOrderRequest struct {
	Symbol                string   `json:"symbol"`
	OrderIdList           []int64  `json:"orderIdList"`
	OrigClientOrderIdList []string `json:"origClientOrderIdList"`
	RecvWindow            int64    `json:"recvWindow"`
	Timestamp             int64    `json:"timestamp"`
}

type BatchCancelOrderReply []CancelOrderReply

type BatchCancelAllOrderRequest struct {
	Symbol     string `json:"symbol"`
	RecvWindow int64  `json:"recvWindow"`
	Timestamp  int64  `json:"timestamp"`
}

type BatchCancelAllOrderReply struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

type FetchOneOrderRequest struct {
	Symbol            string `json:"symbol"`
	OrderId           int64  `json:"orderId"`
	OrigClientOrderId string `json:"origClientOrderId"`
	RecvWindow        int64  `json:"recvWindow"`
	Timestamp         int64  `json:"timestamp"`
}

type OrderInfo struct {
	AvgPrice                string `json:"avgPrice"`
	ClientOrderId           string `json:"clientOrderId"`
	CumQuote                string `json:"cumQuote"`
	ExecutedQty             string `json:"executedQty"`
	OrderId                 int    `json:"orderId"`
	OrigQty                 string `json:"origQty"`
	OrigType                string `json:"origType"`
	Price                   string `json:"price"`
	ReduceOnly              bool   `json:"reduceOnly"`
	Side                    string `json:"side"`
	PositionSide            string `json:"positionSide"`
	Status                  string `json:"status"`
	StopPrice               string `json:"stopPrice"`
	ClosePosition           bool   `json:"closePosition"`
	Symbol                  string `json:"symbol"`
	Time                    int64  `json:"time"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	ActivatePrice           string `json:"activatePrice"`
	PriceRate               string `json:"priceRate"`
	UpdateTime              int64  `json:"updateTime"`
	WorkingType             string `json:"workingType"`
	PriceProtect            bool   `json:"priceProtect"`
	PriceMatch              string `json:"priceMatch"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	GoodTillDate            int    `json:"goodTillDate"`
}

type FetchOneOrderReply OrderInfo

type FetchOpenOrderRequest struct {
	Symbol     string `json:"symbol"`
	RecvWindow int64  `json:"recvWindow,omitempty"`
	Timestamp  int64  `json:"timestamp"`
}

type FetchOpenOrderReply []OrderInfo
