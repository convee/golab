package main

import (
	"encoding/json"
	"fmt"
)

func js2map() {
	var m map[string]int32
	js := `{"aaa": 123, "bbb": 456}`
	json.Unmarshal([]byte(js), &m)
	fmt.Println(m)
}

func js2array() {
	var str []string
	js := `["abc", "123"]`
	json.Unmarshal([]byte(js), &str)
	fmt.Println(str)
}

type P struct {
	Name string `json:"name"`
}

func js2struct() {
	js := `{"name": "abc"}`
	p := P{}
	json.Unmarshal([]byte(js), &p)
	fmt.Println(p)
}

type loginGuestReq struct {
	Did     string           `json:"did"`
	Unionid string           `json:"unionid"`
	Imei    string           `json:"imei"`
	Imsi    string           `json:"imsi"`
	Mac     string           `json:"mac"`
	Psystem string           `json:"psystem"`
	Pmodel  string           `json:"pmodel"`
	Nick    string           `json:"nick"`
	Gender  bool             `json:"gender"`
	Others  map[string]int32 `json:"others"`
}

func js2structWithinMap() {
	req := loginGuestReq{}
	js := `{"did":"70f36b40-d1bc-44e3-9791-9615f3633b45","unionid":"ifunbow","imei":"55932fx23","imsi":"dkdkd12xf9222dk","psystem":"ios","pmode":"iphone5","nick":"一二三四五六七八九十","gender":true,"others":{"character":2}}`
	json.Unmarshal([]byte(js), &req)
	fmt.Println(req)
	res := loginGuestReq{
		Did:     "didxx",
		Unionid: "unxx",
		Imei:    "imei",
		Imsi:    "imsi",
		Psystem: "psystem",
		Pmodel:  "pmodel",
		Nick:    "nick",
		Gender:  true,
		Others:  map[string]int32{"character": 2},
	}
	s, _ := json.Marshal(res)
	fmt.Printf("%s\n", s)
}

func main() {
	js2map()
	js2array()
	js2struct()
	js2structWithinMap()
}
