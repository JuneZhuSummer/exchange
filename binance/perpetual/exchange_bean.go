package perpetual

type RateLimit struct {
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	RateLimitType string `json:"rateLimitType"`
}

type Asset struct {
	Asset             string `json:"asset"`
	MarginAvailable   bool   `json:"marginAvailable"`
	AutoAssetExchange string `json:"autoAssetExchange"`
}

type Filter struct {
	FilterType        string `json:"filterType"`
	MaxPrice          string `json:"maxPrice,omitempty"`
	MinPrice          string `json:"minPrice,omitempty"`
	TickSize          string `json:"tickSize,omitempty"`
	MaxQty            string `json:"maxQty,omitempty"`
	MinQty            string `json:"minQty,omitempty"`
	StepSize          string `json:"stepSize,omitempty"`
	Limit             int    `json:"limit,omitempty"`
	Notional          string `json:"notional,omitempty"`
	MultiplierUp      string `json:"multiplierUp,omitempty"`
	MultiplierDown    string `json:"multiplierDown,omitempty"`
	MultiplierDecimal string `json:"multiplierDecimal,omitempty"`
}

type Symbol struct {
	Symbol                string   `json:"symbol"`
	Pair                  string   `json:"pair"`
	ContractType          string   `json:"contractType"`
	DeliveryDate          int64    `json:"deliveryDate"`
	OnboardDate           int64    `json:"onboardDate"`
	Status                string   `json:"status"`
	MaintMarginPercent    string   `json:"maintMarginPercent"`
	RequiredMarginPercent string   `json:"requiredMarginPercent"`
	BaseAsset             string   `json:"baseAsset"`
	QuoteAsset            string   `json:"quoteAsset"`
	MarginAsset           string   `json:"marginAsset"`
	PricePrecision        int      `json:"pricePrecision"`
	QuantityPrecision     int      `json:"quantityPrecision"`
	BaseAssetPrecision    int      `json:"baseAssetPrecision"`
	QuotePrecision        int      `json:"quotePrecision"`
	UnderlyingType        string   `json:"underlyingType"`
	UnderlyingSubType     []string `json:"underlyingSubType"`
	SettlePlan            int      `json:"settlePlan"`
	TriggerProtect        string   `json:"triggerProtect"`
	Filters               []Filter `json:"filters"`
	OrderType             []string `json:"OrderType"`
	TimeInForce           []string `json:"timeInForce"`
	LiquidationFee        string   `json:"liquidationFee"`
	MarketTakeBound       string   `json:"marketTakeBound"`
}

type ExchangeInfoReply struct {
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	RateLimits      []RateLimit   `json:"rateLimits"`
	ServerTime      int64         `json:"serverTime"`
	Assets          []Asset       `json:"assets"`
	Symbols         []Symbol      `json:"symbols"`
	Timezone        string        `json:"timezone"`
}
