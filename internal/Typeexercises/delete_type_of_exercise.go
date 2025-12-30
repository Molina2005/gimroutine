package typeexercises

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerExercises) DeleteTypeOfExercises(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Metodo no permitido", 404)
		return
	}
	IdParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(IdParam)
	if err != nil {
		http.Error(w, "Id invalido", 400)
		return
	}
	if err := h.service.ServiceDeleteTypeOfExercise(id); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte("tipo ejercicio eliminado con exito"))
}
