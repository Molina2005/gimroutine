package exercises

import (
	"encoding/json"
	"net/http"
)

func (h *HandlerExercises) HandlerCreationExercises(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metodo no permitido", 404)
		return
	}
	var input struct {
		IdTypeOfExercise   int    `json:"idTypeOfExercise"`
		NameTypeOfExercise string `json:"nameTypeOfExercise"`
		Description        string `json:"description"`
		Image              string `json:"image"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON invalido", 400)
		return
	}
	if err := h.service.ServiceCreationExercises(
		input.IdTypeOfExercise,
		input.NameTypeOfExercise,
		input.Description,
		input.Image); err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	w.Write([]byte("ejercicio creado correctamente"))
}
