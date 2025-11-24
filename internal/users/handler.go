package users

// Recibe el servicio con la logica interna
type Handler struct {
	service *Service
}

// Creacion Handler para la conexion entre usuario y servidor
func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}
