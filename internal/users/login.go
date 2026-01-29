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
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON invalido", 400)
		return
	}
	if _, err := h.service.ServiceLogin(input.Email, input.Password); err != nil {
		http.Error(w, "Credenciales incorrectas", 400)
		return
	}
	w.Write([]byte("Inicio sesion correcto"))
}
