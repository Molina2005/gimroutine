package regional

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerRegional) HandlerUpdateRegional(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	// Recibe la peticion con el id a buscar y hace la conversion a int
	IdParam := chi.URLParam(r, "id")
	IdConv, err := strconv.Atoi(IdParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id_invalido",
		})
		return
	}
	// struct para guardar los campos que vienen de peticion json
	var inputUpdate struct {
		Name      string `json:"name"`
		Id_client int    `json:"id_client"`
		Id_state  int    `json:"id_state"`
	}
	if err := json.NewDecoder(r.Body).Decode(&inputUpdate); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "json_invalido",
		})
		return
	}
	if err := h.service.ServiceUpdateRegional(
		IdConv,
		inputUpdate.Id_client,
		inputUpdate.Id_state, inputUpdate.Name,
	); err != nil {
		if errors.Is(err, ErrRegionalDoesNotExists) {
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
