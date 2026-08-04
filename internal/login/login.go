package login

import (
	"encoding/json"
	httpmiddlewares "modulo/internal/http_middlewares"
	"net/http"
)

func (h *HandlerLogin) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	// Header que hace que siempre la respuesta a enviar sea de tipo json
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// Decodifica el JSON recibido en el body de la petición
	// y lo convierte en la estructura input.
	// Si el JSON es inválido, responde con  StatusBadRequest.
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "json_invalido",
		})
		return
	}
	// Se llama al servicio para validar las credenciales.
	// Si ocurre un error, se responde con estado 401 y un mensaje en formato JSON.

	user, err := h.S.ServiceLogin(input.Email, input.Password)
	if err != nil {
		// WriteHeader: Guarda el codigo de estado de la respuesta
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Credenciales invalidas",
		})
		return
	}
	// Se llama funcion para poder obtener el token JWT de cada usuario al ingresar
	token, err := httpmiddlewares.GenerateJWT(user.Id, user.Role)
	if err != nil {
		http.Error(w, "error al generar token", http.StatusUnauthorized)
	}
	// Envio de informacion al frontend
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login exitoso",
		"token":   token,
		// Pasa informacion de si es cliente o usuario
		"type": user.Type,
		"role": user.Role,
	})
}
