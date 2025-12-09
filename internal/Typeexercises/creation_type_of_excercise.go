package typeexercises

import (
	"encoding/json"
	"net/http"
)

func (h *HandlerExercises) HandlerCreationTypeOfExercise(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "metodo no permitido", 404)
		return
	}
	var inputTypeOfExercise struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&inputTypeOfExercise); err != nil {
		http.Error(w, "JSON invalido", 400)
		return
	}
	if err := h.service.ServiceCreationTypeOfExercise(inputTypeOfExercise.Name, inputTypeOfExercise.Description); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte("Tipo de ejercicio creado correctamente"))
}
