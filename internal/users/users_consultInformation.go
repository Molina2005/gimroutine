package users

import (
	"net/http"
)

func (re *Repository) ConsultUserInformation(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Metodo no permitido", 404)
	}

	// Seguir añadiendo logica de consulta

}
