package users

import (
	"errors"
	"fmt"
	"modulo/internal/utils"
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
	// Contraseña incriptada
	HashPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("error al hashear contraseña")
	}
	// Validaion de se pemiten campos vacios
	if name == "" || email == "" || age <= 0 || weight <= 0 || height <= 0 || HashPassword == "" {
		return errors.New("todos los campos son obligatorios")
	}
	return r.repo.InsertUser(name, email, age, weight, height, HashPassword)
}
