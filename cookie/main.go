package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	CLIENT_COOKIE_ID = "clientCookieId"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}

//设置过期时间，cookie会保存到硬盘，可以用于自动登录，记住密码
func setCookieExpired(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    CLIENT_COOKIE_ID,
		Value:   "123",
		Expires: time.Now().Add(180 * time.Second),
	}
	http.SetCookie(w, &cookie)
	fmt.Println("setCookie")
}

//保存cookie在浏览器 重启浏览器 消失
func setCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  CLIENT_COOKIE_ID,
		Value: "456",
	})
}

func getCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(CLIENT_COOKIE_ID)
	if err != nil {
		fmt.Println("get cookie failed,err:", err)
		return
	}
	fmt.Println("cookie:", cookie)
}
func main() {
	http.HandleFunc("/", getCookieHandler)
	http.HandleFunc("/set", setCookieExpired)
	http.HandleFunc("/set2", setCookie)
	http.ListenAndServe("127.0.0.1:5001", nil)
}
