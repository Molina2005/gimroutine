package servicepoints

// Recibe el servicio con la logica interna
type HandlerServicePoints struct {
	service *ServiceServicePoints
}

// Creacion Handler para la conexion entre usuario y servidor
func NewHandler(s *ServiceServicePoints) *HandlerServicePoints {
	return &HandlerServicePoints{service: s}
}
