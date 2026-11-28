package webhook

import (
	"encoding/json"
	"fmt"
	"log"
	"modulo/internal/webhook/models"

	"net/http"
)

func (h *HandlerWebHook) NotifyChangesVitals(w http.ResponseWriter, r *http.Request) {
	var data models.WebHookVitals
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println("error de json", err)
	}
	log.Println("datawebhook data", data)
	w.WriteHeader(200)
}
