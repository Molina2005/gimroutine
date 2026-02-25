package users

import (
	"encoding/json"
	"net/http"
)

func (h *HandlerUsers) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// Decodifica el JSON recibido en el body de la petición
	// y lo convierte en la estructura input.
	// Si el JSON es inválido, responde con 400 Bad Request.
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}
	// Se llama al servicio para validar las credenciales.
	// Si ocurre un error, se responde con estado 401 y un mensaje en formato JSON.
	if _, err := h.service.ServiceLogin(input.Email, input.Password); err != nil {
		// WriteHeader: Guarda el codigo de estado de la respuesta
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Credenciales invalidas"})
		return
	}
	// Envio a frontend login exitoso
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login exitoso"})
}
