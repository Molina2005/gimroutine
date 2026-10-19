package regional

import (
	"errors"
	"strings"
)

// Recibe la conexion real db de Repository
type ServiceRegional struct {
	repo *RepositoryRegional
}

// Creacion nuevo servicio el cual va a guardar toda la logica interna
func NewService(r *RepositoryRegional) *ServiceRegional {
	return &ServiceRegional{repo: r}
}

// Variables que guarda el error de existencia de usuario
var ErrRegionalAlreadyExists = errors.New("regional ya existe en el sistema")
var ErrRegionalDoesNotExists = errors.New("regional no existe en el sistema")

// Creacion de usuario y requirimientos a seguir
func (s *ServiceRegional) ServiceCreateRegional(id_client, id_user_state int, name string) error {
	// Validacion de existencia por nombre de regional
	exists, err := s.repo.QueryRegionalExistsXName(name)
	if err != nil {
		return err
	}
	if exists {
		return ErrRegionalAlreadyExists
	}
	// Nombre en Mayuscula para evitar usuarios duplicados en la creacion
	LowerName := strings.ToLower(name)
	// Validaion de se pemiten campos vacios
	if LowerName == "" || id_client == 0 || id_user_state == 0 {
		return errors.New("todos los campos son obligatorios")
	}
	return s.repo.QueryInsertRegional(id_client, id_user_state, name)
}

// Servicio Consulta datos de regional
func (s *ServiceRegional) ServiceQueryRegional(id_regional int) (*DataRegional, error) {
	// Validacion de existencia por id de regional
	exists, err := s.repo.QueryRegionalExistsXId(id_regional)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrRegionalDoesNotExists
	}
	return s.repo.QueryViewRegionalInfomation(id_regional)
}

// Servicio consultar todas las regionales existentes
func (s *ServiceRegional) ServiceQueryAllRegional() ([]AllRegionales, error) {
	return s.repo.QueryRegionalInformation()
}

// Servicio con logica para validar y enviar informacion de actualizacion de regionales
func (s *ServiceRegional) ServiceUpdateRegional(id_regional, id_client, id_state int, name string) error {
	// Validaciones correspondientes para saber si regional existe
	existsId, errId := s.repo.QueryRegionalExistsXId(id_regional)
	if errId != nil {
		return errId
	}
	existsName, errName := s.repo.QueryRegionalExistsXName(name)
	if errName != nil {
		return errName
	}
	if !existsId {
		return ErrRegionalDoesNotExists
	}
	if existsName {
		return ErrRegionalAlreadyExists
	}
	return s.repo.QueryRegionalUpdate(id_regional, id_client, id_state, name)
}

// Servicio para eliminar regional
func (s *ServiceRegional) ServiceRegionalDelete(id_regional int) error {
	existsId, err := s.repo.QueryRegionalExistsXId(id_regional)
	if err != nil {
		return err
	}
	if !existsId {
		return ErrRegionalDoesNotExists
	}
	return s.repo.QueryRegionalDelete(id_regional)
}

// Servicio para buscador
func (s *ServiceRegional) ServiceRegionalSearch(search string) ([]RegionalSearch, error) {
	return s.repo.QueryRegionalSearch(search)
}
