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

// Variables que guarda el error de existencia de usuario
var ErrExerciseAlreadyExists = errors.New("ejercicio ya existe en el sistema")
var ErrExerciseDoesNotExists = errors.New("ejercicio no existe en el sistema")

// Servicio creacion de ejercicios
func (s *ServiceExercises) ServiceCreationExercises(IdTypeOfExercise int, nameTypeOfExercise, description, image string) error {
	// Verificacion de existencia de ejercicio
	Exist, err := s.Repo.ExistsExercise(nameTypeOfExercise)
	if err != nil {
		return err
	}
	if Exist {
		return ErrExerciseAlreadyExists
	}
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
	Exist, err := s.Repo.ExistsExerciseId(idExercise)
	if err != nil {
		return nil, err
	}
	if !Exist {
		return nil, ErrExerciseDoesNotExists
	}
	return s.Repo.QueryExercises(
		idExercise)
}

// Servicio actualizacion de ejercicios
func (s *ServiceExercises) ServiceUpdateExercises(IdExercise, IdTypeOfExercise int, nameTypeOfExercise, description, image string) error {
	Exist, err := s.Repo.ExistsExerciseId(IdExercise)
	if err != nil {
		return err
	}
	if !Exist {
		return ErrExerciseDoesNotExists
	}
	return s.Repo.QueryUpdateExercise(
		IdExercise,
		IdTypeOfExercise,
		nameTypeOfExercise,
		description,
		image)
}

// Servicio eliminacion de ejercicio
func (s *ServiceExercises) ServiceDeleteExercises(IdExercise int) error {
	Exist, err := s.Repo.ExistsExerciseId(IdExercise)
	if err != nil {
		return err
	}
	if !Exist {
		return ErrExerciseDoesNotExists
	}
	return s.Repo.QueryDeleteExercise(
		IdExercise)
}
