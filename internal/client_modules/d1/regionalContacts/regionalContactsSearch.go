package regionalcontacts

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Handler para manejar peticiones de buscador dinamico de contactos regional
func (h *HandlerRegionalContacts) HandlerRegionalContactsSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	// Obtiene la busqueda que el usuario ingresa y la pasa al service
	// Obtiene id de la regional para poder procesar informacion de solo pertenecientes a esa regional
	search := chi.URLParam(r, "search")
	idRegional := chi.URLParam(r, "idRegional")
	id, err := strconv.Atoi(idRegional)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id_invalido",
		})
	}
	data, err := h.service.ServiceRegionalContactsSeach(search, id)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(data)
}
