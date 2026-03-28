package clients

import "errors"

type ServiceClients struct {
	Repo *RepositoryClients
}

func NewService(r *RepositoryClients) *ServiceClients {
	return &ServiceClients{Repo: r}
}

// Error de existencia del cliente
var ErrClientAlreadyExists = errors.New("Cliente ya existe en el sistema")

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
	if nameClient == "" || document == "" || gmail == "" || phone == "" || password == "" || state == "" {
		return errors.New("todos los campos son obligatorios")
	}
	return s.Repo.QueryCreateClient(nameClient, document, gmail, phone, password, state)
}
