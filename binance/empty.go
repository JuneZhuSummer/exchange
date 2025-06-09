package binance

type Empty struct{}

type Reply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
