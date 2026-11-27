package webkooks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Modelo para obtener los datos de la tabla del evento
type Vital struct {
	IdVital        int       `json:"idvital"`
	IdServicePoint int       `json:"idservicepoint"`
	IdUser         int       `json:"iduser"`
	Incident       int       `json:"incident"`
	Equipment      string    `json:"equipment"`
	Description    string    `json:"description"`
	Oc             string    `json:"oc"`
	EarrIng        string    `json:"earring"`
	EntryDate      time.Time `json:"entrydate"`
	ExecutionDate  time.Time `json:"executiondate"`
	Status         int       `json:"status"`
}

// Modelo para recibir el evento
type data struct {
	Type      string `json:"type"`
	Table     string `json:"table"`
	Schema    string `json:"schema"`
	Record    Vital  `json:"record"`
	OldRecord any    `json:"oldrecord"`
}

func NotifyChangesVitals(w http.ResponseWriter, r *http.Request) {
	var webHook data
	if err := json.NewDecoder(r.Body).Decode(&webHook); err != nil {
		fmt.Println("error de json", err)
	}
	fmt.Println("datawebhook", webHook)
	w.WriteHeader(200)
}
