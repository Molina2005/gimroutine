package typeexercises

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerExercises) HandlerUpdateInfoTypeOfExercises(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Metodo no permitido", 404)
		return
	}
	IdParam := chi.URLParam(r, "id")
	Id, err := strconv.Atoi(IdParam)
	if err != nil {
		http.Error(w, "ID invalido", 400)
		return
	}
	var inputUpdate struct {
		NameTypeOfExercise string `json:"nameTypeOfExercise"`
		Description        string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&inputUpdate); err != nil {
		http.Error(w, "JSON invalido", 400)
		return
	}
	if err := h.service.ServiceUpdateInfoTypeOfExercise(
		Id,
		inputUpdate.NameTypeOfExercise,
		inputUpdate.Description,
	); err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	w.Write([]byte("Tipo de ejercicio actualizado correctamente"))
}
