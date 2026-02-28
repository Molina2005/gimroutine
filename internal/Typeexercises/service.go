package typeexercises

import (
	"errors"
	"modulo/internal/models"
)

type ServiceTypeOfExercises struct {
	repo *RepositoryTypeOfExercises
}

func NewService(r *RepositoryTypeOfExercises) *ServiceTypeOfExercises {
	return &ServiceTypeOfExercises{repo: r}
}

// Requirimientos para la creacion de tipo de ejercicios
func (s *ServiceTypeOfExercises) ServiceCreationTypeOfExercise(nameTypeOfExercise, description string) error {
	if nameTypeOfExercise == "" || description == "" {
		return errors.New("todos los campos son obligatorios")
	}
	// Si no hay error ejecuta la consulta de Repository
	return s.repo.QueryCreateExerciseType(nameTypeOfExercise, description)
}

// Servicio para consultar tipos de ejercicios
func (s *ServiceTypeOfExercises) ServiceQueryTypeOfExercise(IdTypeOfExercise int) (*models.TypeOfExercises, error) {
	return s.repo.QueryTypeOfExercises(IdTypeOfExercise)
}

// Servicio actualizacion informacion tipos de ejercicios
func (s *ServiceTypeOfExercises) ServiceUpdateInfoTypeOfExercise(IdTypeOfExercise int, nameTypeOfExercise, description string) error {
	return s.repo.QueryUpdateTypeOfExercises(IdTypeOfExercise, nameTypeOfExercise, description)
}

// Servicio eliminacion tipos de ejercicios
func (s *ServiceTypeOfExercises) ServiceDeleteTypeOfExercise(IdTypeOfExercise int) error {
	return s.repo.QueryDeleteTypeOfExercises(IdTypeOfExercise)
}
