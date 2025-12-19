package exercises

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryExercises struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *RepositoryExercises {
	return &RepositoryExercises{db: db}
}

// Creacion de ejercicios
func (r *RepositoryExercises) QueryCreateExercises(IdTypeOfExercise int, nameTypeOfExercise, description, image string) error {
	ctx := context.Background()
	query := `INSERT INTO ejercicios (id_tipo_ejercicio, nombre, descripcion, imagen) 
	VALUES ($1,$2,$3,$4)`
	_, err := r.db.Exec(ctx, query, IdTypeOfExercise, nameTypeOfExercise, description, image)
	if err != nil {
		return err
	}
	return nil
}
