package clients

type ServiceClients struct {
	Repo *RepositoryClients
}

func NewService(r *RepositoryClients) *ServiceClients {
	return &ServiceClients{Repo: r}
}
