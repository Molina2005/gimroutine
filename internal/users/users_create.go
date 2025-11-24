package users

import (
	"encoding/json"
	"net/http"
	"time"
)

// Conexion entre usuario y servidor con http por medio de POST
func (h *Handler) CreateUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metodo no permitido", 404)
		return
	}

	// struct para guardar los campos que vienen de peticion json
	var input struct {
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Age       int       `json:"age"`
		Weight    int16     `json:"weight"`
		Height    float64   `json:"height"`
		EntryDate time.Time `json:"entryDate"`
		Password  string    `json:"password"`
	}

	// r.Body : contiene lo que envía el usuario (JSON)
	// Decode(&input) : toma json y convierte a struct para saber qué campos esperar y cómo guardarlos
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON invalido", 400)
		return
	}

	// se pasa la funcion de creacion de usuario con la informacion que esta en input
	if err := h.service.CreatetUser(input.Name, input.Email, input.Age, input.Weight, input.Height, input.EntryDate, input.Password); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Respuesta al usuario de (usuario creado)
	w.Write([]byte("Usuario creado"))
}
