package plans

import (
	"errors"
	"modulo/internal/models"
	"strings"
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

// Creacion de plan y requirimientos a seguir
func (s *ServicePlans) ServiceCreatetPlans(name, description string, userMax int) error {
	// Se pasa el nombre siempre a minuscula y evita duplicados
	// Se manejan los espacios y evita duplicado por espacios entre palabras
	LowerCaseName := strings.ToLower(strings.Join(strings.Fields(name), " "))
	// Validacion para saber si plan ya esta en sistema y evitar creacion
	ExistPlan, err := s.repo.QueryPlanExistOfName(LowerCaseName)
	if err != nil {
		return err
	}
	if ExistPlan {
		return ErrUserAlreadyaPlans
	}
	// Validaion de se pemiten campos vacios
	if LowerCaseName == "" || description == "" || userMax == 0 {
		return errors.New("todos los campos son obligatorios")
	}
	return s.repo.QueryInsertPlans(LowerCaseName, description, userMax)
}

// Servicio para llamar todos los planes
func (s *ServicePlans) ServiceAllPlans() ([]models.PlansAll, error) {
	return s.repo.QueryAllPlans()
}
