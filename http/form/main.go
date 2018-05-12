package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const form = `
<html>
<body>
<form method="post">
<input type="text" name="in1">
<input type="text" name="in2">
<input type="submit">
</form>
</body>
</html>
`

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello,world</h1>")
	panic("test1")
}

func FormServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		r.ParseForm()
		io.WriteString(w, r.FormValue("in1"))
		io.WriteString(w, "\r\n")
		io.WriteString(w, r.FormValue("in2"))
	}
}

func logPanics(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", r.RemoteAddr, x)
			}
		}()
		handler(w, r)
	}
}
func main() {
	http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":5002", nil); err != nil {
		fmt.Printf("server start error:%v", err)
		return
	}
}
