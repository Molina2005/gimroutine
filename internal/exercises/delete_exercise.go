package exercises

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerExercises) HandlerDeleteExercise(w http.ResponseWriter, r *http.Request) {
	IdParam := chi.URLParam(r, "id")
	Id, err := strconv.Atoi(IdParam)
	if err != nil {
		http.Error(w, "ID invalido", 400)
	}
	if r.Method != http.MethodDelete {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	if err := h.service.ServiceDeleteExercises(Id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Ejercicio eliminado correctamente"))
}
