package clients

type HandlerClients struct {
	service *ServiceClients
}

func NewHandler(s *ServiceClients) *HandlerClients {
	return &HandlerClients{service: s}
}
