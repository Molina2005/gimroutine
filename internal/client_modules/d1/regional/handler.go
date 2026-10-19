package regional

// Recibe el servicio con la logica interna
type HandlerRegional struct {
	service *ServiceRegional
}

// Creacion Handler para la conexion entre usuario y servidor
func NewHandler(s *ServiceRegional) *HandlerRegional {
	return &HandlerRegional{service: s}
}
