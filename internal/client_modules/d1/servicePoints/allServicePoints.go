package servicepoints

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerServicePoints) HandlerAllServicePoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	// Se obtiene id segun id que este en url y asi poder pasarlo al service
	idParam := chi.URLParam(r, "idRegional")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id_invalido",
		})
		return
	}
	ServicePoints, err := h.service.ServiceAllServicePoints(id)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(ServicePoints)
}
