package plans

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// Conexion entre usuario y servidor con http por medio de POST
func (h *HandlerPlans) HandlerCreatePlans(w http.ResponseWriter, r *http.Request) {
	// Header que hace que siempre la respuesta a enviar sea de tipo json
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	// struct para guardar los campos que vienen de peticion json
	var input struct {
		Name           string    `json:"name"`
		Description    string    `json:"description"`
		Price          int       `json:"price"`
		DurationMonths int       `json:"Durationmonths"`
		MaxUser        int       `json:"maxuser"`
		ExpirationDate time.Time `json:"expirationdate"`
	}
	// r.Body : contiene lo que envía el usuario (JSON)
	// Decode(&input) : toma json y convierte a struct para saber qué campos esperar y cómo guardarlos
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "json_invalido",
		})
		return
	}
	// se pasa la funcion de creacion de plan con la informacion que esta en input
	if err := h.S.repo.QueryInsertPlans(input.Name, input.Description, input.Price, input.DurationMonths, input.MaxUser, input.ExpirationDate); err != nil {
		if errors.Is(err, ErrUserAlreadyaPlans) {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(map[string]string{
			// Culaquier otro error que venga del service
			"error": err.Error(),
		})
		return
	}
	// Respuesta al usuario de (plan creado) en tipo json
	json.NewEncoder(w).Encode(map[string]string{
		"message": "plan_creado_correctamente",
	})
}
