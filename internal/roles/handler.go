package roles

type HandlerRoles struct {
	s *ServiceRoles
}

func NewHandler(h *ServiceRoles) *HandlerRoles {
	return &HandlerRoles{s: h}
}
