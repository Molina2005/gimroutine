package clients

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerClients) HandlerClientsSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	// Obtiene la busqueda que el cliente ingresa y la pasa al service
	search := chi.URLParam(r, "search")
	data, err := h.service.ServiceClientSearch(search)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(data)
}
