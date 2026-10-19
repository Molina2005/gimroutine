package roles

import (
	"errors"
)

type ServiceRoles struct {
	repo *RepoRoles
}

func NewService(s *RepoRoles) *ServiceRoles {
	return &ServiceRoles{repo: s}
}

var ErrRolAlreadyExists = errors.New("Rol ya existe en el sistema")

// Servicio consulta de todos los roles
func (s *ServiceRoles) ServiceAllRoles() ([]Rol, error) {
	return s.repo.QueryAllRoles()
}
