package vitalincidents

// Recibe la conexion real db de Repository
type ServiceVitalIncidents struct {
	repo *RepositoryVitalIncidents
}

// Creacion nuevo servicio el cual va a guardar toda la logica interna
func NewService(r *RepositoryVitalIncidents) *ServiceVitalIncidents {
	return &ServiceVitalIncidents{repo: r}
}

// Servicio validar existencia de incidencia por id

// Servicio validar existencia de incidencia por numero de incidencia

// Servicio creacion de incidencias

// Servicio para obtener incidencia especifica

// Servicio obtener todas las incidencias segun la regional

// Servicio actualizacion informacion de incidencia

// Servicio eliminacion de incidencia

// Servicio buscador dinamico
