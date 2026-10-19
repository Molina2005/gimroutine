package roles

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepoRoles struct {
	db *pgxpool.Pool
}

func NewRepository(r *pgxpool.Pool) *RepoRoles {
	return &RepoRoles{db: r}
}

// Funcion para manejar existencia del rol
func (r *RepoRoles) QueryExistsRol(post string) (bool, error) {
	ctx := context.Background()
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM rol WHERE cargo=$1)`
	err := r.db.QueryRow(ctx, query, post).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// Funcion consultar todos los roles
func (r *RepoRoles) QueryAllRoles() ([]Rol, error) {
	ctx := context.Background()
	query := `SELECT id_rol, cargo FROM rol`
	data, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	var dataRoles []Rol
	for data.Next() {
		var dataRol Rol
		if err := data.Scan(&dataRol.Id_rol, &dataRol.Post); err != nil {
			return nil, err
		}
		dataRoles = append(dataRoles, dataRol)
	}
	return dataRoles, nil
}
