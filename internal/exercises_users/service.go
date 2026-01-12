package exercisesusers

type ServiceExercisesUsers struct {
	repo *RepositoryExercisesUsers
}

func NewService(r *RepositoryExercisesUsers) *ServiceExercisesUsers {
	return &ServiceExercisesUsers{repo: r}
}
