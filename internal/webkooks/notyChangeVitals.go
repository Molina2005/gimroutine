package webkooks

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type data struct {
	Type      string `json:"type"`
	Table     string `json:"table"`
	Schema    string `json:"schema"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	OldRecord any    `json:"oldrecord"`
}

func NotifyChangesVitals(w http.ResponseWriter, r *http.Request) {
	var webHook data
	if err := json.NewDecoder(r.Body).Decode(&webHook); err != nil {
		fmt.Println("error de json")
	}
	fmt.Println("datawebhook", webHook)
	w.WriteHeader(200)
}
