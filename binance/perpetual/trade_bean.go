package perpetual

type TradesRequest struct {
	Symbol     string `json:"symbol"`
	OrderId    int64  `json:"orderId,omitempty"`
	StartTime  int64  `json:"startTime,omitempty"`
	EndTime    int64  `json:"endTime,omitempty"`
	FromId     int64  `json:"fromId,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	RecvWindow int64  `json:"recvWindow,omitempty"`
	Timestamp  int64  `json:"timestamp"`
}

type Trade struct {
	Buyer           bool   `json:"buyer"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Id              int    `json:"id"`
	Maker           bool   `json:"maker"`
	OrderId         int    `json:"orderId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	QuoteQty        string `json:"quoteQty"`
	RealizedPnl     string `json:"realizedPnl"`
	Side            string `json:"side"`
	PositionSide    string `json:"positionSide"`
	Symbol          string `json:"symbol"`
	Time            int64  `json:"time"`
}

type TradesReply []Trade
