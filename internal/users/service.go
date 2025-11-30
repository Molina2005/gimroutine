package users

import (
	"errors"
	// "time"
)

// Recibe la conexion real db de Repository
type Service struct {
	repo *Repository
}

// Creacion nuevo servicio el cual va a guardar toda la logica interna
func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

// Creacion de usuario y requirimientos a seguir
func (r *Service) CreatetUser(name, email string, age int, weight int16, height float64, password string) error {

	if name == "" || email == "" || age <= 0 || weight <= 0 || height <= 0 || password == "" {
		return errors.New("todos los campos son obligatorios")
	}
	return r.repo.InsertUser(name, email, age, weight, height, password)
}
