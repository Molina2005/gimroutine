package users

import (
	"encoding/json"
	"net/http"
)

func (h *HandlerUsers) HandlerConsultAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	Users, err := h.service.ServiceConsultAllUsers()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(Users)
}
