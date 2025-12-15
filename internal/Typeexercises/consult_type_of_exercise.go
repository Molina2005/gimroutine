package typeexercises

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (re *RepositoryExercises) HandlerConsultTypeOfExercises(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Metodo no permitido", 404)
		return
	}
	IdParam := chi.URLParam(r, "id")
	Id, err := strconv.Atoi(IdParam)
	if err != nil {
		http.Error(w, "ID invalido", 400)
		return
	}
	InfoTypeOfExercise, err := re.QueryTypeOfExercises(Id)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	json.NewEncoder(w).Encode(InfoTypeOfExercise)
}
