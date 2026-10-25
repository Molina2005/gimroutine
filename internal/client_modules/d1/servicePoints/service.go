package servicepoints

import "errors"

// Recibe la conexion real db de Repository
type ServiceServicePoints struct {
	repo *RepositoryServicePoints
}

// Creacion nuevo servicio el cual va a guardar toda la logica interna
func NewService(r *RepositoryServicePoints) *ServiceServicePoints {
	return &ServiceServicePoints{repo: r}
}

var ErrServicePointsAlreadyExists = errors.New("Punto de servicio ya existe en el sistema")
var ErrServicePointsDoesNotExists = errors.New("Punto de servicio no existe en el sistema")

// Servicio creacion de puntos servicios
func (s *ServiceServicePoints) ServiceServicePointsCreation(name, addres, city string, idRegional, idState int) error {
	// Funcion para validacion de existencia punto servicio en sistema
	ExistsName, errName := s.repo.ExistsServicePointsOfName(name)
	if errName != nil {
		return errName
	}
	if ExistsName {
		return ErrServicePointsAlreadyExists
	}
	return s.repo.QueryServicePointCreate(name, addres, city, idRegional, idState)
}

// Servicio obtener informacion de todos los puntos de servicios
func (s *ServiceServicePoints) ServiceAllServicePoints(idRegional int) ([]ServicePoints, error) {
	return s.repo.QueryAllServicePoints(idRegional)
}

// Servicio para poder consultar a punto servicio especifico
func (s *ServiceServicePoints) ServicePointInformation(idServicePoint int) (*ServicePoint, error) {
	ExistsId, errId := s.repo.ExistsServicePointsOfId(idServicePoint)
	if errId != nil {
		return nil, errId
	}
	if !ExistsId {
		return nil, ErrServicePointsAlreadyExists
	}
	return s.repo.QueryServicePointInformation(idServicePoint)
}
