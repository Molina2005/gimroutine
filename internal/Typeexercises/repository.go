package typeexercises

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

// Verificacion existencia de ejercicios por nombre
func (r *RepositoryExercises) QueryExerciseExistsName(nameTypeOfExercise string) (bool, error) {
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

// Verificacion existencia de ejercicios por id
func (r *RepositoryExercises) QueryExerciseExistsId(IdTypeOfExercise int) (bool, error) {
	var ExistsTypeOfExercise bool
	ctx := context.Background()
	r.db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM tipo_ejercicios WHERE id_tipo_ejercicio = $1)", IdTypeOfExercise).Scan(&ExistsTypeOfExercise)
	if !ExistsTypeOfExercise {
		return false, nil
	}
	return ExistsTypeOfExercise, nil
}

// Creacion de tipos de ejercicios
func (r *RepositoryExercises) QueryCreateExerciseType(nameTypeOfExercise, description string) error {
	Exists, err := r.QueryExerciseExistsName(nameTypeOfExercise)
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

// Consulta de tipos de ejercicios
func (r *RepositoryExercises) QueryTypeOfExercises(IdTypeOfExercise int) (*models.TypeOfExercises, error) {
	Exists, err := r.QueryExerciseExistsId(IdTypeOfExercise)
	if err != nil {
		return nil, err
	}
	if !Exists {
		return nil, errors.New("tipo de ejercicio no existe en el sistema")
	}
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
func (r *RepositoryExercises) QueryUpdateTypeOfExercises(IdTypeOfExercise int, nameTypeOfExercise, description string) error {
	ctx := context.Background()
	Exists, err := r.QueryExerciseExistsId(IdTypeOfExercise)
	if err != nil {
		return err
	}
	if !Exists {
		return errors.New("tipo de ejercicio no existe en el sistema")
	}
	query := `UPDATE tipo_ejercicios SET nombre = $1, descripcion = $2
	WHERE id_tipo_ejercicio = $3`
	_, err = r.db.Exec(ctx, query, nameTypeOfExercise, description, IdTypeOfExercise)
	if err != nil {
		return err
	}
	return nil
}

// Eliminar tipos de ejercicios
func (r *RepositoryExercises) QueryDeleteTypeOfExercises(idTypeOfExercise int) error {
	ctx := context.Background()
	Exists, err := r.QueryExerciseExistsId(idTypeOfExercise)
	if err != nil {
		return err
	}
	if !Exists {
		return errors.New("tipo de ejercicio no existe en el sistema")
	}
	query := `DELETE FROM tipo_ejercicios WHERE id_tipo_ejercicio = $1`
	_, err = r.db.Exec(ctx, query, idTypeOfExercise)
	if err != nil {
		return err
	}
	return nil
}
