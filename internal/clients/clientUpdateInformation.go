package clients

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerClients) HandlerUpdateClients(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	IdParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(IdParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id_invalido",
		})
		return
	}
	var input struct {
		Name     string `json:"name"`
		Document string `json:"document"`
		Gmail    string `json:"gmail"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
		State    string `json:"state"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "json_invalido",
		})
		return
	}
	if err := h.service.ServiceUpdateClient(
		id,
		input.Name,
		input.Document,
		input.Gmail,
		input.Phone,
		input.Password,
		input.State); err != nil {
		if errors.Is(err, ErrClientDoesNotExists) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "informacion_actualizada_correctamente",
	})
}
