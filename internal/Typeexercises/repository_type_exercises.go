package typeexercises

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryExercises struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *RepositoryExercises {
	return &RepositoryExercises{db: db}
}

// Verificacion existencia de ejercicios por nombre
func (r *RepositoryExercises) QueryExerciseExists(nameTypeOfExercise string) (bool, error) {
	var ExistsTypeOfExercise bool
	ctx := context.Background()
	r.db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM tipo_ejercicios WHERE nombre = $1)", nameTypeOfExercise).Scan(&ExistsTypeOfExercise)
	// Si no existe genera el error
	if !ExistsTypeOfExercise {
		return false, nil
	}
	// si existe retorna su existencia
	return ExistsTypeOfExercise, nil
}

// Creacion de tipos de ejercicios
func (r *RepositoryExercises) QueryCreateExerciseType(nameTypeOfExercise, description string) error {
	Exists, err := r.QueryExerciseExists(nameTypeOfExercise)
	if err != nil {
		return err
	}
	if Exists {
		return errors.New("tipo de ejercicio ya existe en el sistema")
	}
	if !Exists {
		ctx := context.Background()
		query := `INSERT INTO tipo_ejercicios (nombre, descripcion) 
		VALUES ($1,$2)`
		_, err := r.db.Exec(ctx, query, nameTypeOfExercise, description)
		return err
	}
	return nil
}
