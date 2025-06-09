package binance

import "fmt"

type Error struct {
	Method       string
	Url          string
	RawE         error
	Msg          string
	ResponseBody string
}

func (e *Error) Error() string {
	if e.RawE == nil {
		return e.ResponseBody
	}
	return fmt.Sprintf("method:%s, url:%s, raw_err:%v, msg:%s", e.Method, e.Url, e.RawE, e.Msg)
}
