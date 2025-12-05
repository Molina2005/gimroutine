package users

// Recibe el servicio con la logica interna
type HandlerUsers struct {
	service *ServiceUsers
}

// Creacion Handler para la conexion entre usuario y servidor
func NewHandler(s *ServiceUsers) *HandlerUsers {
	return &HandlerUsers{service: s}
}
