package users

import (
	"errors"
	"fmt"
	"modulo/internal/utils"
)

// Recibe la conexion real db de Repository
type ServiceUsers struct {
	repo *RepositoryUsers
}

// Creacion nuevo servicio el cual va a guardar toda la logica interna
func NewService(r *RepositoryUsers) *ServiceUsers {
	return &ServiceUsers{repo: r}
}

// Creacion de usuario y requirimientos a seguir
func (r *ServiceUsers) CreatetUser(name, email string, password string) error {
	// Contraseña incriptada
	HashPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("error al hashear contraseña")
	}
	// Validaion de se pemiten campos vacios
	if name == "" || email == "" || HashPassword == "" {
		return errors.New("todos los campos son obligatorios")
	}
	return r.repo.InsertUser(name, email, HashPassword)
}

// Requerimientos actualizacion informacion usuario
func (r *ServiceUsers) UpdateUserInformation(id_usuarios int, name, email string, password string) error {
	HashPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("error al hashear contraseña")
	}
	return r.repo.UpdateUser(id_usuarios, name, email, HashPassword)
}

// Requerimientos eliminacion usuario
func (r *ServiceUsers) DeleteUser(id_usuarios int) error {
	return r.repo.DeleteUser(id_usuarios)
}
