package clients

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryClients struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *RepositoryClients {
	return &RepositoryClients{db: db}
}

// Crear fncion valdar exixtencia por id
// Crear fncion valdar exixtencia por correo
// Crear fncion valdar exixtencia por documento

// consulta para creacion de clientes
func (r *RepositoryClients) QueryCreateClient(name, document, gmail, phone, password, state string) error {
	ctx := context.Background()
	query := `INSERT INTO clientes (nombre,documento,correo,telefono,contrasena,estado) 
				VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(ctx, query, name, document, gmail, phone, password, state)
	if err != nil {
		return err
	}
	return nil
}
