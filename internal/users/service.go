package users

import (
	"errors"
	"fmt"
	"modulo/internal/models"
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
func (r *ServiceUsers) ServiceCreatetUser(name, email string, password string) error {
	// Contraseña incriptada
	HashPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("error al hashear contraseña")
	}
	// Validaion de se pemiten campos vacios
	if name == "" || email == "" || HashPassword == "" {
		return errors.New("todos los campos son obligatorios")
	}
	return r.repo.QueryInsertUser(name, email, HashPassword)
}

// Requerimientos actualizacion informacion usuario
func (r *ServiceUsers) ServiceUpdateUserInformation(id_usuarios int, name, email string, password string) error {
	HashPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("error al hashear contraseña")
	}
	return r.repo.QueryUpdateUser(id_usuarios, name, email, HashPassword)
}

// Requerimientos eliminacion usuario
func (r *ServiceUsers) ServiceDeleteUser(id_usuarios int) error {
	return r.repo.QueryDeleteUser(id_usuarios)
}

// Requerimientos login
func (r *ServiceUsers) ServiceLogin(email, password string) (*models.Login, error) {
	if email == "" || password == "" {
		return nil, errors.New("email y contraseña son obligatorios")
	}
	return r.repo.QueryLogin(email, password)
}
