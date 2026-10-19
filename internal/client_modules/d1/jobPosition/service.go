package jobposition

type ServiceJobPosition struct {
	repo *RepositoryJobPosition
}

func NewService(s *RepositoryJobPosition) *ServiceJobPosition {
	return &ServiceJobPosition{repo: s}
}

// Funcion para logica interna y funcionalidades para traer todos los cargos de trabajo
func (s *ServiceJobPosition) ServiceAllJobPosition() ([]JobPositions, error) {
	return s.repo.QueryAllJobPosition()
}
