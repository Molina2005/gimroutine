package regionalcontacts

type HandlerRegionalContacts struct {
	service *ServiceRegionalContacts
}

func NewHandler(s *ServiceRegionalContacts) *HandlerRegionalContacts {
	return &HandlerRegionalContacts{service: s}
}
