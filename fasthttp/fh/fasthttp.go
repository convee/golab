package golab

import (
	"crypto/tls"
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {
	client := fasthttp.Client{TLSConfig: &tls.Config{InsecureSkipVerify: true}}
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("url")
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	resp := fasthttp.AcquireResponse()
	req.SetBodyString(`{"body": "json_str"}`)   //设置请求参数
	req.SetBody([]byte(`{"body": "sjon_str"}`)) //设置[]byte
	if err := client.Do(req, resp); err != nil {
		fmt.Printf("loan list fail to do request. appID=%s. [err=%v]\n", h["X-PPD-APPID"], err)
		continue
	}
	b := resp.Body()
	if resp.StatusCode() != fasthttp.StatusOK {
		fmt.Printf("loan list failed code=%d. [err=%v]\n", resp.StatusCode(), string(b))
		continue
	}
}
