package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

type Person struct {
	Title string
	Name  string
	Age   int
}

func initTemplate(filename string) (err error) {
	template, err := t.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	p := Person{
		Title: "我的个人博客",
		Name:  "convee",
		Age:   22,
	}
	if err := t.Execute(w, p); err != nil {
		fmt.Println("There is an error:", err.Error())
	}
}

func main() {
	filename := "c:/Users/conve/go/src/golab/http/template/index.html"
	initTemplate(filename)
	//输出到浏览器
	http.HandleFunc("/", userInfo)
	http.ListenAndServe(":5003", nil)

}
