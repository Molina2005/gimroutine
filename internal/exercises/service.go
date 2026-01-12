package exercises

import "errors"

type ServiceExercises struct {
	Repo *RepositoryExercises
}

func NewService(r *RepositoryExercises) *ServiceExercises {
	return &ServiceExercises{Repo: r}
}

// Servicio creacion de ejercicios
func (s *ServiceExercises) ServiceCreationExercises(IdTypeOfExercise int, nameTypeOfExercise, description, image string) error {
	if IdTypeOfExercise <= 0 || nameTypeOfExercise == "" || description == "" || image == "" {
		return errors.New("todos los campos son obligatorios")
	}
	return s.Repo.QueryCreateExercises(
		IdTypeOfExercise,
		nameTypeOfExercise,
		description,
		image)
}

// Servicio actualizacion de ejercicios
func (s *ServiceExercises) ServiceUpdateExercises(IdExercise, IdTypeOfExercise int, nameTypeOfExercise, description, image string) error {
	return s.Repo.QueryUpdateExercise(
		IdExercise,
		IdTypeOfExercise,
		nameTypeOfExercise,
		description,
		image)
}

func (s *ServiceExercises) ServiceDeleteExercises(IdExercise int) error {
	return s.Repo.QueryDeleteExercise(
		IdExercise)
}
