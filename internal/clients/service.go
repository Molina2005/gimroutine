package clients

import (
	"errors"
	"modulo/internal/login"
)

type ServiceClients struct {
	Repo      *RepositoryClients
	RepoLogin *login.RepoLogin
}

func NewService(r *RepositoryClients, rl *login.RepoLogin) *ServiceClients {
	return &ServiceClients{Repo: r, RepoLogin: rl}
}

// Error de existencia del cliente
var ErrClientAlreadyExists = errors.New("Cliente ya existe en el sistema")
var ErrClientDoesNotExists = errors.New("Cliente no existe en el sistema")

// Servicios internos para creacion de clientes
func (s *ServiceClients) ServiceCreateClient(name, phone, gmail, nit string, id_state int) error {
	existGmailClient, err := s.RepoLogin.QueryClientExistsByGmail(gmail)
	if err != nil {
		return err
	}
	// Validacion de existencia de correo
	existGmailUser, err := s.RepoLogin.QueryuserExistsXEmail(gmail)
	if err != nil {
		return err
	}
	// Si existe muestra error de existencia
	if existGmailClient || existGmailUser {
		return ErrClientAlreadyExists
	}
	// Validacion de existencia de documento
	ExistsDoc, err := s.Repo.QueryClientExistsByDocument(gmail)
	if err != nil {
		return err
	}
	// Si existe muestra error de existencia
	if ExistsDoc {
		return ErrClientAlreadyExists
	}
	return s.Repo.QueryCreateClient(name, phone, gmail, nit, id_state)
}

// Servicio consultar cliente
func (s *ServiceClients) ServiceQueryClient(id_client int) (*Client, error) {
	Exist, err := s.Repo.QueryClientExistsById(id_client)
	if err != nil {
		return nil, err
	}
	if !Exist {
		return nil, ErrClientDoesNotExists
	}
	return s.Repo.QueryViewClientInfomation(id_client)
}

// Servicio consultar todos los clientes
func (s *ServiceClients) ServiceConsultAllClient() ([]Clients, error) {
	return s.Repo.QueryClientInformation()
}

// Servicio actualizar informacion clientes
func (s *ServiceClients) ServiceUpdateClient(id_client int, name, phone, nit, gmail string, id_estado_usuario int) error {
	Exist, err := s.Repo.QueryClientExistsById(id_client)
	if err != nil {
		return err
	}
	if !Exist {
		return ErrClientDoesNotExists
	}
	return s.Repo.QueryUpdateClient(id_client, name, gmail, phone, nit, id_estado_usuario)
}

// Requerimientos eliminacion cliente
func (s *ServiceClients) ServiceDeleteClient(id_client int) error {
	Exist, err := s.Repo.QueryClientExistsById(id_client)
	if err != nil {
		return err
	}
	if !Exist {
		return ErrClientDoesNotExists
	}
	return s.Repo.QueryDeleteClient(id_client)
}

// Servicio para buscar cliente por nombre documento o correo
func (s *ServiceClients) ServiceClientSearch(dataClientSearch string) ([]SearchClients, error) {
	return s.Repo.QuerySearchClients(dataClientSearch)
}
