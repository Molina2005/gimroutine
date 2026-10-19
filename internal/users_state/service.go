package usersstate

type ServiceUserState struct {
	repo *RepositoryUserState
}

func NewService(s *RepositoryUserState) *ServiceUserState {
	return &ServiceUserState{repo: s}
}

// Funcion para logica interna y funcionalidades para traer todos los status incidencias
func (s *ServiceUserState) ServiceAllIncidenceState() ([]UserStatus, error) {
	return s.repo.QueryAllUserStates()
}
