package vitalincidents

// Recibe el servicio con la logica interna
type HandlerVitalIncidents struct {
	service *ServiceVitalIncidents
}

// Creacion Handler para la conexion entre usuario y servidor
func NewHandler(s *ServiceVitalIncidents) *HandlerVitalIncidents {
	return &HandlerVitalIncidents{service: s}
}
