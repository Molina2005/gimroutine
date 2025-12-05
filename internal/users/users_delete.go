package users

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerUsers) DeleteUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Metodo no permitido", 404)
		return
	}
	IdParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(IdParam)
	if err != nil {
		http.Error(w, "Id invalido", 400)
		return
	}
	if err := h.service.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte("usuario eliminado con exito"))
}
