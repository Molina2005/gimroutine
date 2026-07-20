package clients

import (
	"errors"
	"modulo/internal/models"
	"modulo/internal/utils"
)

type ServiceClients struct {
	Repo *RepositoryClients
}

func NewService(r *RepositoryClients) *ServiceClients {
	return &ServiceClients{Repo: r}
}

// Error de existencia del cliente
var ErrClientAlreadyExists = errors.New("Cliente ya existe en el sistema")
var ErrClientDoesNotExists = errors.New("Cliente no existe en el sistema")

func (s *ServiceClients) ServiceCreateClient(nameClient, document, gmail, phone, password, state string) error {
	// Validacion de existencia de cliente por correo
	ExistsGmail, err := s.Repo.QueryClientExistsByGmail(gmail)
	if err != nil {
		return err
	}
	// Validacion de existencia de documento
	ExistsDoc, err := s.Repo.QueryClientExistsByGmail(document)
	if err != nil {
		return err
	}
	// Si existe muestra error de existencia
	if ExistsDoc {
		return ErrClientAlreadyExists
	}
	if ExistsGmail {
		return ErrClientAlreadyExists
	}
	// Hasheo de contraseña
	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	return s.Repo.QueryCreateClient(nameClient, document, gmail, phone, hash, state)
}

func (s *ServiceClients) ServiceConsultAllClient() ([]models.Client, error) {
	return s.Repo.QueryClientInformation()
}

func (s *ServiceClients) ServiceUpdateClient(id_client int, name, document, gmail, phone, password, state string) error {
	Exist, err := s.Repo.QueryClientExistsById(id_client)
	if err != nil {
		return err
	}
	if !Exist {
		return ErrClientDoesNotExists
	}
	if password == "" {
		return s.Repo.QueryUpdateClientNoPassword(id_client, name, document, gmail, phone, state)
	}
	HashPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	return s.Repo.QueryUpdateClient(id_client, name, document, gmail, phone, HashPassword, state)
}
