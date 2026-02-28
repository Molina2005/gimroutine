package exercises

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerExercises) HandlerConsultInformationExercise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	Idparam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(Idparam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id_invalido",
		})
		return
	}
	InfoExercise, err := h.service.ServiceQueryExercises(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	// Creacion nuevo encoder para poder enviarle la informacion al usuario en dato json
	json.NewEncoder(w).Encode(map[string]any{
		"exercise": InfoExercise,
	})
}
