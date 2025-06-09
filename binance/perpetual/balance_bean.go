package perpetual

type BalanceRequest struct {
	RecvWindow int64 `json:"recvWindow,omitempty"`
	Timestamp  int64 `json:"timestamp"`
}

type BalanceMeta struct {
	AccountAlias       string `json:"accountAlias"`
	Asset              string `json:"asset"`
	Balance            string `json:"balance"`
	CrossWalletBalance string `json:"crossWalletBalance"`
	CrossUnPnl         string `json:"crossUnPnl"`
	AvailableBalance   string `json:"availableBalance"`
	MaxWithdrawAmount  string `json:"maxWithdrawAmount"`
	MarginAvailable    bool   `json:"marginAvailable"`
	UpdateTime         int64  `json:"updateTime"`
}

type BalanceReply []BalanceMeta

type AccountRequest struct {
	RecvWindow int64 `json:"recvWindow,omitempty"`
	Timestamp  int64 `json:"timestamp"`
}

type AccountAsset struct {
	Asset                  string `json:"asset"`
	WalletBalance          string `json:"walletBalance"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	MarginBalance          string `json:"marginBalance"`
	MaintMargin            string `json:"maintMargin"`
	InitialMargin          string `json:"initialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	AvailableBalance       string `json:"availableBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	UpdateTime             int64  `json:"updateTime"`
}

type AccountPosition struct {
	Symbol           string `json:"symbol"`
	PositionSide     string `json:"positionSide"`
	PositionAmt      string `json:"positionAmt"`
	UnrealizedProfit string `json:"unrealizedProfit"`
	IsolatedMargin   string `json:"isolatedMargin"`
	Notional         string `json:"notional"`
	IsolatedWallet   string `json:"isolatedWallet"`
	InitialMargin    string `json:"initialMargin"`
	MaintMargin      string `json:"maintMargin"`
	UpdateTime       int    `json:"updateTime"`
}

type AccountReply struct {
	TotalInitialMargin          string            `json:"totalInitialMargin"`
	TotalMaintMargin            string            `json:"totalMaintMargin"`
	TotalWalletBalance          string            `json:"totalWalletBalance"`
	TotalUnrealizedProfit       string            `json:"totalUnrealizedProfit"`
	TotalMarginBalance          string            `json:"totalMarginBalance"`
	TotalPositionInitialMargin  string            `json:"totalPositionInitialMargin"`
	TotalOpenOrderInitialMargin string            `json:"totalOpenOrderInitialMargin"`
	TotalCrossWalletBalance     string            `json:"totalCrossWalletBalance"`
	TotalCrossUnPnl             string            `json:"totalCrossUnPnl"`
	AvailableBalance            string            `json:"availableBalance"`
	MaxWithdrawAmount           string            `json:"maxWithdrawAmount"`
	Assets                      []AccountAsset    `json:"assets"`
	Positions                   []AccountPosition `json:"positions"`
}
