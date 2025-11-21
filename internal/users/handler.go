package users

import (
	"net/http"
)

// Recibe el servicio con la logica interna
type Handler struct {
	service *Service
}

// Creacion nuevo Handler para la conexion entre usuario y servidor
func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

// Conexion entre usuario y servidor con http por medio de POST
func (h *Handler) CreateUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metodo no permitido", 404)
		return
	}

	// var input struct {
	// 	Name      string    `json:"name"`
	// 	Email     string    `json:"email"`
	// 	Age       int       `json:"age"`
	// 	Weight    int16     `json:"weight"`
	// 	Height    float64   `json:"height"`
	// 	EntryDate time.Time `json:"entryDate"`
	// 	Password  string    `json:"password"`
	// }
}
