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
func (s *ServicePlans) ServiceCreatetPlans(name, description string, userMax, Months, Price int) error {
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
	if LowerCaseName == "" || description == "" || userMax == 0 || Months == 0 || Price == 0 {
		return errors.New("todos los campos son obligatorios")
	}
	// Recibe datos basicos de plan y pasa a consulta db para poder retornar el id ya creado
	// y pasarlo a QueryInsertPrecio
	Id, err := s.repo.QueryInsertPlans(LowerCaseName, description, userMax)
	// Se pasa el id del plan creado y se crea su regitro con meses y precio en tabla
	return s.repo.QueryInsertPrecio(Id, Months, Price)
}

// Servicio para llamar todos los planes
func (s *ServicePlans) ServiceAllPlans() ([]models.PlansAll, error) {
	return s.repo.QueryAllPlans()
}
