package typeexercises

import (
	"errors"
)

type ServiceExercises struct {
	repo *RepositoryExercises
}

func NewService(r *RepositoryExercises) *ServiceExercises {
	return &ServiceExercises{repo: r}
}

// Requirimientos para la creacion de tipo de ejercicios
func (s *ServiceExercises) ServiceCreationTypeOfExercise(nameTypeOfExercise, description string) error {
	if nameTypeOfExercise == "" || description == "" {
		return errors.New("todos los campos son obligatorios")
	}
	// Si no hay error ejecuta la consulta de Repository
	return s.repo.QueryCreateExerciseType(nameTypeOfExercise, description)
}

// Servicio actualizacion informacion tipos de ejercicios
func (s *ServiceExercises) ServiceUpdateInfoTypeOfExercise(IdTypeOfExercise int, nameTypeOfExercise, description string) error {
	return s.repo.QueryUpdateTypeOfExercises(IdTypeOfExercise, nameTypeOfExercise, description)
}

// Servicio eliminacion tipos de ejercicios
func (s *ServiceExercises) ServiceDeleteTypeOfExercise(IdTypeOfExercise int) error {
	return s.repo.QueryDeleteTypeOfExercises(IdTypeOfExercise)
}
