package exercises

type ServiceExercises struct {
	Repo *RepositoryExercises
}

func NewService(r *RepositoryExercises) *ServiceExercises {
	return &ServiceExercises{Repo: r}
}
