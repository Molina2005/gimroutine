package regionalcontacts

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Handler para manejar peticiones para obtener todos los contactos regional
func (h *HandlerRegionalContacts) HandlerRegionalContacts(w http.ResponseWriter, r *http.Request) {
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
	InfoRegionalContact, err := h.service.ServiceRegionalContactsInformation(Id)
	if err != nil {
		if errors.Is(err, ErrRegionalContactsDoesNotExists) {
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
		"regionaContact": InfoRegionalContact,
	})

}
