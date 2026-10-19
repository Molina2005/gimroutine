package vitalincidents

// Recibe la conexion real db de Repository
type ServiceVitalIncidents struct {
	repo *RepositoryVitalIncidents
}

// Creacion nuevo servicio el cual va a guardar toda la logica interna
func NewService(r *RepositoryVitalIncidents) *ServiceVitalIncidents {
	return &ServiceVitalIncidents{repo: r}
}
