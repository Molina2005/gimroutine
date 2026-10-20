package regionalcontacts

import (
	"errors"
)

type ServiceRegionalContacts struct {
	Repo *RepositoryRegionalContacts
}

func NewService(r *RepositoryRegionalContacts) *ServiceRegionalContacts {
	return &ServiceRegionalContacts{Repo: r}
}

var ErrContactAlreadyExists = errors.New("Contacto ya existe en el sistema")
var ErrRegionalContactsDoesNotExists = errors.New("Contacto no existe en el sistema")

// Servicio para crear contactos regional
func (s *ServiceRegionalContacts) ServiceRegionalContactsCreate(name, lastName, gmail, phone string,
	id_job_position, id_regional, id_user_state int) error {
	ExistsName, errName := s.Repo.QueryRegionalContactsExitsByName(name)
	if errName != nil {
		return errName
	}
	ExistsGmail, errGmail := s.Repo.QueryRegionalContactsExitsByGmail(gmail)
	if errGmail != nil {
		return errGmail
	}
	if ExistsName || ExistsGmail {
		return ErrContactAlreadyExists
	}
	return s.Repo.QueryRegionalContactsCreate(name, lastName, gmail, phone, id_job_position, id_regional, id_user_state)
}

// Servicio para consultar contacto de regional especifico
func (s *ServiceRegionalContacts) ServiceRegionalContactsInformation(idRegionalContact int) (*regionalContact, error) {
	ExistsId, err := s.Repo.QueryRegionalContactsExitsById(idRegionalContact)
	if err != nil {
		return nil, err
	}
	if !ExistsId {
		return nil, ErrRegionalContactsDoesNotExists
	}
	return s.Repo.QueryRegionalContactsInformation(idRegionalContact)
}

// Servicio para consultar todos los contactos regional
func (s *ServiceRegionalContacts) ServiceAllRegionalContacts(id_regional int) ([]RegionalesContacts, error) {
	return s.Repo.QueryAllRegionalContacts(id_regional)
}

// Servicio para actualizar informacion del contacto
func (s *ServiceRegionalContacts) ServiceRegionalContactsUpdate(
	name, lastName, gmail, phone string,
	id_job_position, id_regional, id_user_state, idregionalContact int) error {
	existsId, errId := s.Repo.QueryRegionalContactsExitsById(idregionalContact)
	if errId != nil {
		return errId
	}
	if !existsId {
		return ErrRegionalContactsDoesNotExists
	}
	return s.Repo.QueryRegionalContactsUpdate(name, lastName, gmail, phone,
		id_job_position, id_regional, id_user_state, idregionalContact)
}

// Servicio eliminar contacto regional
func (s *ServiceRegionalContacts) ServiceRegionalContactsDelete(idRegionalContact int) error {
	return s.Repo.QueryRegionalContactsDelete(idRegionalContact)
}

// Servicio buscador de contactos regional
func (s *ServiceRegionalContacts) ServiceRegionalContactsSeach(search string, idRegional int) ([]searchRegionalContacts, error) {
	return s.Repo.QueryRegionalSeach(search, idRegional)
}
