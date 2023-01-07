package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	hurl "net/url"
	"sort"
	"strings"
)

func main() {
	params := map[string]string{
		"a": "1",
		"b": "2",
	}
	result := getSign("https://www.convee.cn", "post", "scret", params)

	fmt.Println(result)
}

func getSign(url string, method string, scret string, params map[string]string) string {
	var keyArr []string
	for k, v := range params {
		if k == "auth_signature" || strings.HasPrefix(v, "@") {
			continue
		}
		keyArr = append(keyArr, k)
	}
	sort.Strings(keyArr)
	var final []string
	for _, value := range keyArr {
		vv, _ := params[value]
		final = append(final, fmt.Sprintf("%s=%s", hurl.QueryEscape(value), hurl.QueryEscape(vv)))
	}
	str := strings.ToUpper(method) + "&" + hurl.QueryEscape(strings.ToLower(url)) + "&" + hurl.QueryEscape(strings.Join(final, "&"))
	fmt.Println(str)
	return GetSignature(str, scret)
}

func GetSignature(input, key string) string {
	key_for_sign := []byte(key)
	h := hmac.New(sha1.New, key_for_sign)
	h.Write([]byte(input))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
