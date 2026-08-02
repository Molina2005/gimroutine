package clients

import (
	"errors"
	"modulo/internal/login"
	"modulo/internal/models"
	"modulo/internal/utils"
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

func (s *ServiceClients) ServiceCreateClient(nameClient, document, gmail, phone, password, state string) error {
	existGmailClient, err := s.RepoLogin.QueryClientExistsByGmail(gmail)
	if err != nil {
		return err
	}
	existGmailUser, err := s.RepoLogin.QueryuserExistsXEmail(gmail)
	if err != nil {
		return err
	}
	if existGmailClient || existGmailUser {
		return ErrClientAlreadyExists
	}
	// Validacion de existencia de documento
	ExistsDoc, err := s.Repo.QueryClientExistsByDocument(document)
	if err != nil {
		return err
	}
	// Si existe muestra error de existencia
	if ExistsDoc {
		return ErrClientAlreadyExists
	}
	// Hasheo de contraseña
	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	return s.Repo.QueryCreateClient(nameClient, document, gmail, phone, hash, state)
}

// Servicio consultar cliente
func (s *ServiceClients) ServiceQueryClient(id_client int) (*models.Client, error) {
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
func (s *ServiceClients) ServiceConsultAllClient() ([]models.Clients, error) {
	return s.Repo.QueryClientInformation()
}

// Servicio actualizar informacion clientes
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
func (s *ServiceClients) ServiceClientSearch(dataClientSearch string) ([]models.SearchClients, error) {
	return s.Repo.QuerySearchClients(dataClientSearch)
}
