package regional

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Conexion entre usuario y servidor con http por medio de POST
func (h *HandlerRegional) HandlerCreateRegional(w http.ResponseWriter, r *http.Request) {
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
		Id_client     int    `json:"id_client"`
		Name          string `json:"name"`
		Id_user_state int    `json:"id_user_state"`
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
	// se pasa la funcion de creacion de regional con la informacion que esta en input
	if err := h.service.ServiceCreateRegional(input.Id_client, input.Id_user_state, input.Name); err != nil {
		if errors.Is(err, ErrRegionalAlreadyExists) {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(map[string]string{
			// Cualquier otro error que venga del service
			"error": err.Error(),
		})
		return
	}
	// Respuesta al usuario de (regional creado) en tipo json
	json.NewEncoder(w).Encode(map[string]string{
		"message": "regional_creado_correctamente",
	})
}
