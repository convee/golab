package main

import (
	"crypto/tls"
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {

	url := "http://creative.iqiyi.com/upload/post"

	client := fasthttp.Client{TLSConfig: &tls.Config{InsecureSkipVerify: true}}
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")

	req.Header.Add("account", "3213")
	req.Header.Add("detail_page", "https%3A%2F%2Frmsandbox.addnewer.com%2F%23%21%2Fcampaign%2Fedit%2F433789%2Fmatter")
	req.Header.Add("ad_id", "26")
	req.Header.Add("ad_sub_type", "1")
	req.Header.Add("ad_type", "14")
	req.Header.Add("click_url", "https%3A%2F%2Fhubdev.addnewer.com%2Fchain%2Flanding.html%3Ft%3DfRI989FJYDu%26bc%3D%24%7BEXT2%7D")
	req.Header.Add("dsp_token", "306322e6437243947b4feb09e9ebd00b")
	req.Header.Add("file_name", "1632454998386250.jpg")
	req.Header.Add("is_pmp", "1")
	req.Header.Add("platform", "2")
	req.Header.Add("title", "title")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp := fasthttp.AcquireResponse()
	//req.SetBodyString(`{"body": "json_str"}`)   //设置请求参数
	//req.SetBody([]byte(`{"body": "sjon_str"}`)) //设置[]byte
	if err := client.Do(req, resp); err != nil {
		fmt.Printf("loan list fail to do request. appID=%s. [err=%v]\n", err)
		return
	}
	b := resp.Body()
	if resp.StatusCode() != fasthttp.StatusOK {
		fmt.Printf("loan list failed code=%d. [err=%v]\n", resp.StatusCode(), string(b))
		return
	}
	fmt.Println(string(b))
}
