package exercises

import (
	"context"
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

// Verificacion existencia de ejericicio por id
func (r *RepositoryExercises) ExistsExerciseId(IdExercise int) (bool, error) {
	var ExistsExercise bool
	ctx := context.Background()
	err := r.db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM ejercicios WHERE id_ejercicio = $1)", IdExercise).Scan(&ExistsExercise)
	if err != nil {
		return false, nil
	}
	return ExistsExercise, nil
}

// Creacion de ejercicios
func (r *RepositoryExercises) QueryCreateExercises(IdTypeOfExercise int, nameExercise, description, image string) error {
	ctx := context.Background()
	query := `INSERT INTO ejercicios (id_tipo_ejercicio, nombre, descripcion, imagen) 
	VALUES ($1,$2,$3,$4)`
	_, err := r.db.Exec(ctx, query, IdTypeOfExercise, nameExercise, description, image)
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

// Actualizar datos de ejercicios con imagen
func (r *RepositoryExercises) QueryUpdateExercise(IdExercise, IdTypeOfExercise int, nameExercise, description, image string) error {
	ctx := context.Background()
	query := `UPDATE ejercicios SET id_tipo_ejercicio = $1, nombre = $2, descripcion = $3, imagen = $4 
	WHERE id_ejercicio = $5`
	if _, err := r.db.Exec(ctx, query, IdTypeOfExercise, nameExercise, description, image, IdExercise); err != nil {
		return err
	}
	return nil
}

// Eliminar ejercicios del sistema
func (r *RepositoryExercises) QueryDeleteExercise(IdExercise int) error {
	ctx := context.Background()
	query := `DELETE FROM ejercicios WHERE id_ejercicio = $1`
	if _, err := r.db.Exec(ctx, query, IdExercise); err != nil {
		return err
	}
	return nil
}
