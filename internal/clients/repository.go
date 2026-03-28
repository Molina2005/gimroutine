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

// Crear funcion validar exixtencia por id
func (r *RepositoryClients) QueryClientExistsById(IdClient int) (bool, error) {
	ctx := context.Background()
	var ExistsId bool
	query := `SELECT * FROM clientes WHERE EXISTS(SELECT 1 FROM clientes WHERE id_cliente = $1)`
	err := r.db.QueryRow(ctx, query, IdClient).Scan(&ExistsId)
	if err != nil {
		return false, err
	}
	return ExistsId, nil
}

// Crear funcion validar existencia por correo
func (r *RepositoryClients) QueryClientExistsByGmail(Gmail string) (bool, error) {
	ctx := context.Background()
	var ExistsGmail bool
	query := `SELECT * FROM clientes WHERE EXISTS(SELECT 1 FROM clientes WHERE correo = $1)`
	err := r.db.QueryRow(ctx, query, Gmail).Scan(&ExistsGmail)
	if err != nil {
		return false, err
	}
	return ExistsGmail, nil
}

// Crear fncion valdar exixtencia por documento
func (r *RepositoryClients) QueryClientExistsByDocument(Document string) (bool, error) {
	ctx := context.Background()
	var ExistsDocument bool
	query := `SELECT * FROM clientes WHERE EXISTS(SELECT 1 FROM clientes WHERE documento = $1)`
	err := r.db.QueryRow(ctx, query, Document).Scan(&ExistsDocument)
	if err != nil {
		return false, err
	}
	return ExistsDocument, nil
}

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
