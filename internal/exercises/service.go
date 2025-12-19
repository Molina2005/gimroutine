package exercises

import "errors"

type ServiceExercises struct {
	Repo *RepositoryExercises
}

func NewService(r *RepositoryExercises) *ServiceExercises {
	return &ServiceExercises{Repo: r}
}

func (s *ServiceExercises) ServiceCreationExercises(IdTypeOfExercise int, nameTypeOfExercise, description, image string) error {
	if IdTypeOfExercise <= 0 || nameTypeOfExercise == "" || description == "" || image == "" {
		return errors.New("todos los campos son obligatorios")
	}
	return s.Repo.QueryCreateExercises(IdTypeOfExercise, nameTypeOfExercise, description, image)
}
