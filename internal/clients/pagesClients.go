package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"modulo/internal/pages"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerClients) HandlerClientsPages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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
	client, err := h.service.ServiceQueryClient(Id)
	if err != nil {
		// Se comparan el err con el error ErrUserDoesNotExists el cual es una variable global que esta en el service
		// Si el error es ErrUserDoesNotExists entra a StatusNotFound, si no se sabe que es un error interno del servidor
		if errors.Is(err, ErrClientDoesNotExists) {
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

	Name := strings.ToLower(client.Name)
	pages, exists := pages.ClientPages[Name]
	if !exists {
		fmt.Print("error en pagina")
		return
	}
	http.ServeFile(w, r, pages)
}
