package incidencestate

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryIncidenceState struct {
	db *pgxpool.Pool
}

func NewRepository(d *pgxpool.Pool) *RepositoryIncidenceState {
	return &RepositoryIncidenceState{db: d}
}

// Funcion para consultar cada uno de los estados de incidencia existentes
func (r *RepositoryIncidenceState) QueryAllIncidenceState() ([]IncidenceStatus, error) {
	ctx := context.Background()
	query := `SELECT id_status_incidencia, nombre FROM status_incidencias`
	data, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var dataAllIncidenceState []IncidenceStatus
	for data.Next() {
		var dataIncidenceState IncidenceStatus
		if err := data.Scan(&dataIncidenceState.Id, &dataIncidenceState.Name); err != nil {
			return nil, err
		}
		dataAllIncidenceState = append(dataAllIncidenceState, dataIncidenceState)
	}
	return dataAllIncidenceState, nil
}
