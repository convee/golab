package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

type TagMeta struct {
	Name    string     `json:"name"`
	Count   uint64     `json:"count,omitempty"`
	Version string     `json:"version,omitempty"`
	Group   []*TagMeta `json:"group,omitempty"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	person := &Person{Name: "convee", Age: 20}
	pjson, _ := json.Marshal(person)
	fmt.Println(string(pjson))
	fmt.Fprintf(w, string(pjson))
}

func Tags(w http.ResponseWriter, r *http.Request) {
	tagMeta := &TagMeta{Name: "convee", Count: 100, Version: "v1.0"}

	pjson, _ := json.Marshal([]*TagMeta{tagMeta})
	fmt.Println(string(pjson))
	fmt.Fprintf(w, string(pjson))
}
func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/tags", Tags)
	err := http.ListenAndServe("0.0.0.0:9091", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
