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

func Hello(w http.ResponseWriter, r *http.Request) {
	person := &Person{Name: "convee", Age: 20}
	pjson, _ := json.Marshal(person)
	fmt.Println(string(pjson))
	fmt.Fprintf(w, string(pjson))
}
func main() {
	http.HandleFunc("/hello", Hello)
	err := http.ListenAndServe("0.0.0.0:5001", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
