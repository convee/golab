package common

import (
	"encoding/json"
	"net/http"
)

func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

type User struct {
	Name string
}

func SendJSON(w http.ResponseWriter, r *http.Request) {
	u := User{
		Name: "zhangsan",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)

}
