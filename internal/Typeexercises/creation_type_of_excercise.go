package typeexercises

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (h *HandlerTypeOfExercises) HandlerCreationTypeOfExercise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	var inputTypeOfExercise struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&inputTypeOfExercise); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "json_invalido",
		})
		return
	}
	if err := h.service.ServiceCreationTypeOfExercise(inputTypeOfExercise.Name, inputTypeOfExercise.Description); err != nil {
		// Comparacion de errores, si el error es que el tipo ejercicio ya existe entra a http.StatusConflict, si no es error interno http.StatusInternalServerError
		if errors.Is(err, ErrTypeOfExerciseAlreadyExists) {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		// Otros errores que pueden venir
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "tipo_ejercicio_creado_correctamente",
	})
}
