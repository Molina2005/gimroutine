package clients

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (h *HandlerClients) HandlerCreationClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	var input struct {
		Name           string `json:"name"`
		Phone          string `json:"phone"`
		Gmail          string `json:"gmail"`
		Nit            string `json:"nit"`
		Id_users_state int    `json:"id_users_state"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "json_invalido",
		})
		return
	}
	if err := h.service.ServiceCreateClient(input.Name, input.Phone, input.Gmail, input.Nit, input.Id_users_state); err != nil {
		if errors.Is(err, ErrClientAlreadyExists) {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "cliente_creado_correctamente",
	})
}
