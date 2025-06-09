package perpetual

type PositionRequest struct {
	Symbol     string `json:"symbol,omitempty"`
	RecvWindow int64  `json:"recvWindow,omitempty"`
	Timestamp  int64  `json:"timestamp"`
}

type PositionMeta struct {
	Symbol                 string `json:"symbol"`
	PositionSide           string `json:"positionSide"`
	PositionAmt            string `json:"positionAmt"`
	EntryPrice             string `json:"entryPrice"`
	BreakEvenPrice         string `json:"breakEvenPrice"`
	MarkPrice              string `json:"markPrice"`
	UnRealizedProfit       string `json:"unRealizedProfit"`
	LiquidationPrice       string `json:"liquidationPrice"`
	IsolatedMargin         string `json:"isolatedMargin"`
	Notional               string `json:"notional"`
	MarginAsset            string `json:"marginAsset"`
	IsolatedWallet         string `json:"isolatedWallet"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	Adl                    int    `json:"adl"`
	BidNotional            string `json:"bidNotional"`
	AskNotional            string `json:"askNotional"`
	UpdateTime             int64  `json:"updateTime"`
}

type PositionReply []PositionMeta

type PositionSideDualRequest struct {
	DualSidePosition bool  `json:"dualSidePosition"`
	RecvWindow       int64 `json:"recvWindow,omitempty"`
	Timestamp        int64 `json:"timestamp"`
}
