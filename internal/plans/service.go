package plans

import (
	"errors"
	"time"
)

// Recibe la conexion real db de Repository
type ServicePlans struct {
	repo *RepositoryPlans
}

// Creacion nuevo servicio el cual va a guardar toda la logica interna
func NewService(r *RepositoryPlans) *ServicePlans {
	return &ServicePlans{repo: r}
}

// Variables que guarda el error de existencia de plan
var ErrUserAlreadyaPlans = errors.New("plan ya existe en el sistema")

// Creacion de usuario y requirimientos a seguir
func (s *ServicePlans) ServiceCreatetUser(name, description string, price, durationMonths, maxUser int, expirationDate time.Time) error {
	// Validacion para saber si plan ya esta en sistema y evitar creacion
	ExistPlan, err := s.repo.QueryPlanExistOfName(name)
	if err != nil {
		return err
	}
	if ExistPlan {
		return ErrUserAlreadyaPlans
	}
	// Validaion de se pemiten campos vacios
	if name == "" || description == "" || price == 0 || durationMonths == 0 || maxUser == 0 || expirationDate.IsZero() {
		return errors.New("todos los campos son obligatorios")
	}
	return s.repo.QueryInsertPlans(name, description, price, durationMonths, maxUser, expirationDate)
}
