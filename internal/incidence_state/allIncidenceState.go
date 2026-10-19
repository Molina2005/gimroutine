package incidencestate

import (
	"encoding/json"
	"net/http"
)

func (h *HandlerIncidenceState) HandlerAllIncidenceState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	UsersStates, err := h.service.ServiceAllIncidenceState()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(UsersStates)
}
