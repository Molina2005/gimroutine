package exercises

import (
	"errors"
	"modulo/internal/models"
)

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

// Servicio consulta de ejercicios
func (s *ServiceExercises) ServiceQueryExercises(idExercise int) (*models.Exercises, error) {
	return s.Repo.QueryExercises(idExercise)
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
