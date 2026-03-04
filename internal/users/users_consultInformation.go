package users

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerUsers) HandlerConsultUserInformation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	// recibe la peticion con el id a buscar y hace la conversion a int
	IdParam := chi.URLParam(r, "id")
	Id, err := strconv.Atoi(IdParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id_invalido",
		})
		return
	}
	// Llamado de funcion al servicio
	InfoUser, err := h.service.ServiceQueryUser(Id)
	if err != nil {
		// Se comparan el err con el error ErrUserDoesNotExists el cual es una variable global que esta en el service
		// Si el error es ErrUserDoesNotExists entra a StatusNotFound, si no se sabe que es un error interno del servidor
		if errors.Is(err, ErrUserDoesNotExists) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		// Cualquier otro error que en envie el service
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	// Creacion nuevo encoder para poder enviarle la informacion al usuario en dato json
	json.NewEncoder(w).Encode(map[string]any{
		"user": InfoUser,
	})
}
