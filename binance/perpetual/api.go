package perpetual

import (
	"context"
	"exchange/binance"
	"fmt"
	"github.com/avast/retry-go/v4"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	BaseURL = "https://fapi.binance.com"
	TestUrl = "https://testnet.binancefuture.com"
)

type API struct {
	endpoint  string
	key       string
	secret    string
	retryOpts []retry.Option
}

func New(endpoint, key, secret string, retryOpts ...retry.Option) *API {
	if retryOpts == nil {
		retryOpts = []retry.Option{
			//重试延迟时间
			retry.Delay(200 * time.Millisecond),
			//重试次数
			retry.Attempts(10),
			//只返回上次重试的错误
			retry.LastErrorOnly(true),
			//固定时间重试
			retry.DelayType(retry.FixedDelay),
		}
	}

	return &API{
		endpoint:  endpoint,
		key:       key,
		secret:    secret,
		retryOpts: retryOpts,
	}
}

func (i *API) get(ctx context.Context, url string, params any) ([]byte, error) {
	origin := i.prepare(params)
	sign := binance.Sign(origin, i.secret)

	url = fmt.Sprintf("%s?%s&signature=%s", url, origin, sign)

	return i.sendRequest(ctx, http.MethodGet, url, "")
}

func (i *API) post(ctx context.Context, url string, params any) ([]byte, error) {
	origin := i.prepare(params)
	sign := binance.Sign(origin, i.secret)

	body := fmt.Sprintf("%s&signature=%s", origin, sign)

	return i.sendRequest(ctx, http.MethodPost, url, body)
}

func (i *API) delete(ctx context.Context, url string, params any) ([]byte, error) {
	origin := i.prepare(params)
	sign := binance.Sign(origin, i.secret)
	url = fmt.Sprintf("%s?%s&signature=%s", url, origin, sign)

	return i.sendRequest(ctx, http.MethodDelete, url, "")
}

func (i *API) put(ctx context.Context, url string, params any) ([]byte, error) {
	origin := i.prepare(params)
	sign := binance.Sign(origin, i.secret)

	body := fmt.Sprintf("%s&signature=%s", origin, sign)

	return i.sendRequest(ctx, http.MethodPut, url, body)
}

func (i *API) sendRequest(ctx context.Context, method, url, body string) ([]byte, error) {
	bts := make([]byte, 0)

	//重试方法体
	retryFunc := func() error {
		resp, err := i.request(ctx, method, url, body)
		if err != nil {
			return err
		}
		bts = resp
		return nil
	}

	err := retry.Do(retryFunc, i.retryOpts...)
	if err != nil {
		return nil, err
	}
	return bts, nil
}

func (i *API) request(ctx context.Context, method, url, body string) ([]byte, error) {
	cli := http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequestWithContext(ctx, method, url, strings.NewReader(body))
	if err != nil {
		e := &binance.Error{
			Method:       method,
			Url:          url,
			RawE:         err,
			Msg:          "new http.Request err",
			ResponseBody: "",
		}
		return nil, e
	}

	req.Header.Add("X-MBX-APIKEY", i.key)

	if method == http.MethodPost && len(body) != 0 {
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
	}

	resp, err := cli.Do(req)
	if err != nil {
		return nil, &binance.Error{
			Method:       method,
			Url:          url,
			RawE:         err,
			Msg:          "send request err",
			ResponseBody: "",
		}
	}
	defer func(Body io.ReadCloser) {
		errB := Body.Close()
		if errB != nil {
			log.Printf("body close err:%+v", errB)
			return
		}
	}(resp.Body)

	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &binance.Error{
			Method:       method,
			Url:          url,
			RawE:         err,
			Msg:          "body parse err",
			ResponseBody: "",
		}
	}
	defer cli.CloseIdleConnections() //关闭空闲连接，防止内存泄漏

	if resp.StatusCode != http.StatusOK {
		return bts, &binance.Error{
			Method:       method,
			Url:          url,
			RawE:         nil,
			Msg:          "response err",
			ResponseBody: string(bts),
		}
	}

	return bts, nil
}

func (i *API) prepare(params any) string {
	m, err := binance.StructToMap(params, binance.TagJson)
	if err != nil {
		return ""
	}
	res := ""
	for key, val := range m {
		if res == "" {
			res = fmt.Sprintf("%s=%v", key, val)
			continue
		}
		res = fmt.Sprintf("%s&%s=%v", res, key, val)
	}
	return res
}
