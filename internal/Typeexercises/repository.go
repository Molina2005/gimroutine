package typeexercises

import (
	"context"
	"modulo/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryTypeOfExercises struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *RepositoryTypeOfExercises {
	return &RepositoryTypeOfExercises{db: db}
}

// Verificacion existencia de ejercicios por nombre
func (r *RepositoryTypeOfExercises) QueryExerciseExistsName(nameTypeOfExercise string) (bool, error) {
	var ExistsTypeOfExercise bool
	ctx := context.Background()
	err := r.db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM tipo_ejercicios WHERE nombre = $1)", nameTypeOfExercise).Scan(&ExistsTypeOfExercise)
	// Si no existe genera el error
	if err != nil {
		return false, nil
	}
	// si existe retorna su existencia
	return ExistsTypeOfExercise, nil
}

// Verificacion existencia de ejercicios por id
func (r *RepositoryTypeOfExercises) QueryExerciseExistsId(IdTypeOfExercise int) (bool, error) {
	var ExistsTypeOfExercise bool
	ctx := context.Background()
	err := r.db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM tipo_ejercicios WHERE id_tipo_ejercicio = $1)", IdTypeOfExercise).Scan(&ExistsTypeOfExercise)
	if err != nil {
		return false, nil
	}
	return ExistsTypeOfExercise, nil
}

// Creacion de tipos de ejercicios
func (r *RepositoryTypeOfExercises) QueryCreateExerciseType(nameTypeOfExercise, description string) error {
	ctx := context.Background()
	query := `INSERT INTO tipo_ejercicios (nombre, descripcion) 
		VALUES ($1,$2)`
	_, err := r.db.Exec(ctx, query, nameTypeOfExercise, description)
	return err
}

// Consulta de tipos de ejercicios
func (r *RepositoryTypeOfExercises) QueryTypeOfExercises(IdTypeOfExercise int) (*models.TypeOfExercises, error) {
	ctx := context.Background()
	query := `SELECT id_tipo_ejercicio, nombre, descripcion, fecha_creacion 
	FROM tipo_ejercicios WHERE id_tipo_ejercicio = $1`
	var DataTypeOfExercise models.TypeOfExercises
	if err := r.db.QueryRow(ctx, query, IdTypeOfExercise).Scan(
		&DataTypeOfExercise.Id,
		&DataTypeOfExercise.Name,
		&DataTypeOfExercise.Description,
		&DataTypeOfExercise.CreationDate,
	); err != nil {
		return nil, err
	}
	return &DataTypeOfExercise, nil
}

// Actualizar informacion de tipos de ejercicios
func (r *RepositoryTypeOfExercises) QueryUpdateTypeOfExercises(IdTypeOfExercise int, nameTypeOfExercise, description string) error {
	ctx := context.Background()
	query := `UPDATE tipo_ejercicios SET nombre = $1, descripcion = $2
	WHERE id_tipo_ejercicio = $3`
	_, err := r.db.Exec(ctx, query, nameTypeOfExercise, description, IdTypeOfExercise)
	if err != nil {
		return err
	}
	return nil
}

// Eliminar tipos de ejercicios
func (r *RepositoryTypeOfExercises) QueryDeleteTypeOfExercises(idTypeOfExercise int) error {
	ctx := context.Background()
	query := `DELETE FROM tipo_ejercicios WHERE id_tipo_ejercicio = $1`
	_, err := r.db.Exec(ctx, query, idTypeOfExercise)
	if err != nil {
		return err
	}
	return nil
}
