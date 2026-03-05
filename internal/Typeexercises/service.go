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

// Variables que guarda el error de existencia de usuario
var ErrTypeOfExerciseAlreadyExists = errors.New("tipo de ejercicio ya existe en el sistema")
var ErrTypeOfExerciseNotDoesNotExists = errors.New("tipo de ejercicio no existe en el sistema")

// Servicio para la creacion de tipo de ejercicios
func (s *ServiceTypeOfExercises) ServiceCreationTypeOfExercise(nameTypeOfExercise, description string) error {
	// Verificacion de existencia de tipo de ejercicio
	Exist, err := s.repo.QueryExerciseExistsName(nameTypeOfExercise)
	if err != nil {
		return err
	}
	// Si existe lanza error de ya existencia
	if Exist {
		return ErrTypeOfExerciseAlreadyExists
	}
	if nameTypeOfExercise == "" || description == "" {
		return errors.New("todos los campos son obligatorios")
	}
	// Si no hay error ejecuta la consulta de Repository
	return s.repo.QueryCreateExerciseType(nameTypeOfExercise, description)
}

// Servicio para consultar tipos de ejercicios
func (s *ServiceTypeOfExercises) ServiceQueryTypeOfExercise(IdTypeOfExercise int) (*models.TypeOfExercises, error) {
	Exist, err := s.repo.QueryExerciseExistsId(IdTypeOfExercise)
	if err != nil {
		return nil, err
	}
	if !Exist {
		return nil, ErrTypeOfExerciseNotDoesNotExists
	}
	return s.repo.QueryTypeOfExercises(IdTypeOfExercise)
}

// Servicio actualizacion informacion tipos de ejercicios
func (s *ServiceTypeOfExercises) ServiceUpdateInfoTypeOfExercise(IdTypeOfExercise int, nameTypeOfExercise, description string) error {
	Exist, err := s.repo.QueryExerciseExistsId(IdTypeOfExercise)
	if err != nil {
		return err
	}
	if !Exist {
		return ErrTypeOfExerciseNotDoesNotExists
	}
	return s.repo.QueryUpdateTypeOfExercises(IdTypeOfExercise, nameTypeOfExercise, description)
}

// Servicio eliminacion tipos de ejercicios
func (s *ServiceTypeOfExercises) ServiceDeleteTypeOfExercise(IdTypeOfExercise int) error {
	Exist, err := s.repo.QueryExerciseExistsId(IdTypeOfExercise)
	if err != nil {
		return err
	}
	if !Exist {
		return ErrTypeOfExerciseNotDoesNotExists
	}
	return s.repo.QueryDeleteTypeOfExercises(IdTypeOfExercise)
}
