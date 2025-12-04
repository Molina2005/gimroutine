package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (re *Repository) ConsultUserInformation(w http.ResponseWriter, r *http.Request) {
	// recibe la peticion con el id a buscar y hace la conversion a int
	IdParam := chi.URLParam(r, "id")
	Id, err := strconv.Atoi(IdParam)
	if err != nil {
		http.Error(w, "Id invalido", 400)
		return
	}
	// Llamado consulta de repository
	InfoUser, err := re.ViewUserInfomation(Id)
	if err != nil {
		// cualquier error de la consulta se va a reflejar aqui con err.Error()
		http.Error(w, err.Error(), 404)
		return
	}
	// Creacion de nuevo encoder json, para poder enviar informacion exacta al usuario
	json.NewEncoder(w).Encode(InfoUser)
}
