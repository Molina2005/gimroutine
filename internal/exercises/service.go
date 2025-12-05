package exercises

import (
	"errors"
	"time"
)

type ServiceExercises struct {
	repo *RepositoryExercises
}

func NewService(r *RepositoryExercises) *ServiceExercises {
	return &ServiceExercises{repo: r}
}

// Requirimientos para la creacion de tipo de ejercicios
func (s *ServiceExercises) CreationTypeOfExercise(nameTypeOfExercise, description string, creationDate time.Time) error {
	if nameTypeOfExercise == "" || description == "" || creationDate.IsZero() {
		return errors.New("todos los campos son obligatorios")
	}
	// Si no hay error ejecuta la consulta de Repository
	return s.repo.CreateExerciseType(nameTypeOfExercise, description, creationDate)
}
