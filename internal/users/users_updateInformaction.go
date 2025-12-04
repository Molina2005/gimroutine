package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) UpdateUsersInformation(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Metodo no permitido", 404)
		return
	}
	// Recibe la peticion con el id a buscar y hace la conversion a int
	IdParam := chi.URLParam(r, "id")
	IdConv, err := strconv.Atoi(IdParam)
	if err != nil {
		http.Error(w, "ID invalido", 404)
	}
	// struct para guardar los campos que vienen de peticion json
	var inputUpdate struct {
		Name     string  `json:"name"`
		Email    string  `json:"email"`
		Age      int     `json:"age"`
		Weight   int16   `json:"weight"`
		Height   float64 `json:"height"`
		Password string  `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&inputUpdate); err != nil {
		http.Error(w, "JSON invalido", 400)
		return
	}
	// se pasa la funcion de creacion de usuario con la informacion que esta en inputUpdate
	if err := h.service.UpdateUserInformation(
		IdConv,
		inputUpdate.Name,
		inputUpdate.Email,
		inputUpdate.Age,
		inputUpdate.Weight,
		inputUpdate.Height,
		inputUpdate.Password,
	); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// Respuesta al usuario
	w.Write([]byte("informacion actualizada"))
}
