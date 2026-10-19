package regionalcontacts

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Handler para manejar peticiones de creacion de contactos regional
func (h *HandlerRegionalContacts) HandlerRegionalContactsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	var input struct {
		Name            string `json:"name"`
		LastName        string `json:"lastname"`
		Gmail           string `json:"gmail"`
		Phone           string `json:"phone"`
		Id_job_position int    `json:"id_job_position"`
		Id_regional     int    `json:"id_regional"`
		Id_users_state  int    `json:"id_user_state"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "json_invalido",
		})
		return
	}
	if err := h.service.ServiceRegionalContactsCreate(input.Name, input.LastName, input.Gmail,
		input.Phone, input.Id_job_position, input.Id_regional, input.Id_users_state); err != nil {
		if errors.Is(err, ErrContactAlreadyExists) {
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
		"message": "contacto_creado_correctamente",
	})

}
