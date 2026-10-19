package jobposition

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryJobPosition struct {
	db *pgxpool.Pool
}

func NewRepository(d *pgxpool.Pool) *RepositoryJobPosition {
	return &RepositoryJobPosition{db: d}
}

// Funcion para consultar cada uno de los cargos de trabajo existentes
func (r *RepositoryJobPosition) QueryAllJobPosition() ([]JobPositions, error) {
	ctx := context.Background()
	query := `SELECT id_cargo, nombre FROM cargos`
	data, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var dataAllJobPosition []JobPositions
	for data.Next() {
		var dataJobPosition JobPositions
		if err := data.Scan(&dataJobPosition.Id, &dataJobPosition.Name); err != nil {
			return nil, err
		}
		dataAllJobPosition = append(dataAllJobPosition, dataJobPosition)
	}
	return dataAllJobPosition, nil
}
