package usersstate

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryUserState struct {
	db *pgxpool.Pool
}

func NewRepository(d *pgxpool.Pool) *RepositoryUserState {
	return &RepositoryUserState{db: d}
}

// Funcion para consultar cada uno de los estados de usuarios existentes
func (r *RepositoryUserState) QueryAllUserStates() ([]UserStatus, error) {
	ctx := context.Background()
	query := `SELECT id_estado_usuario, nombre FROM estados_usuarios`
	data, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var dataAllUserState []UserStatus
	for data.Next() {
		var dataUserState UserStatus
		if err := data.Scan(&dataUserState.Id, &dataUserState.Name); err != nil {
			return nil, err
		}
		dataAllUserState = append(dataAllUserState, dataUserState)
	}
	return dataAllUserState, nil
}
