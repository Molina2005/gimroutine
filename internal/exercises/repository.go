package exercises

import (
	"context"
	"errors"
	"modulo/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryExercises struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *RepositoryExercises {
	return &RepositoryExercises{db: db}
}

// Verificacion existencia de ejericicio por nombre
func (r *RepositoryExercises) ExistsExercise(nameExercise string) (bool, error) {
	var ExistsExercise bool
	ctx := context.Background()
	err := r.db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM ejercicios WHERE nombre = $1)", nameExercise).Scan(&ExistsExercise)
	if err != nil {
		return false, nil
	}
	return ExistsExercise, nil
}

// Creacion de ejercicios
func (r *RepositoryExercises) QueryCreateExercises(IdTypeOfExercise int, nameExercise, description, image string) error {
	Exists, err := r.ExistsExercise(nameExercise)
	if err != nil {
		return err
	}
	if Exists {
		return errors.New("ejercicio ya existe en el sistema")
	}
	ctx := context.Background()
	query := `INSERT INTO ejercicios (id_tipo_ejercicio, nombre, descripcion, imagen) 
	VALUES ($1,$2,$3,$4)`
	_, err = r.db.Exec(ctx, query, IdTypeOfExercise, nameExercise, description, image)
	if err != nil {
		return err
	}
	return nil
}

// Consultar informacion ejercicios
func (r *RepositoryExercises) QueryExercises(idExercise int) (*models.Exercises, error) {
	ctx := context.Background()
	query := `SELECT id_ejercicio, id_tipo_ejercicio, nombre, descripcion, imagen, fecha_creacion 
	FROM ejercicios WHERE id_ejercicio = $1`
	var DataExercise models.Exercises
	if err := r.db.QueryRow(ctx, query, idExercise).Scan(
		&DataExercise.Id,
		&DataExercise.IdTypeOfExercise,
		&DataExercise.Name,
		&DataExercise.Description,
		&DataExercise.Img,
		&DataExercise.CreationDate,
	); err != nil {
		return nil, err
	}
	return &DataExercise, nil
}
