package incidencestate

type ServiceIncidenceState struct {
	repo *RepositoryIncidenceState
}

func NewService(s *RepositoryIncidenceState) *ServiceIncidenceState {
	return &ServiceIncidenceState{repo: s}
}

// Funcion para logica interna y funcionalidades para traer todos los status incidencias
func (s *ServiceIncidenceState) ServiceAllIncidenceState() ([]IncidenceStatus, error) {
	return s.repo.QueryAllIncidenceState()
}
