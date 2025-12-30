package exercises

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (re *RepositoryExercises) HandlerConsultInformationExercise(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metodo no permitido", 404)
		return
	}
	Idparam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(Idparam)
	if err != nil {
		http.Error(w, "ID invalido", 400)
		return
	}
	InfoExercise, err := re.QueryExercises(id)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	// Creacion nuevo encoder para poder enviarle la informacion al usuario en dato json
	json.NewEncoder(w).Encode(InfoExercise)
}
