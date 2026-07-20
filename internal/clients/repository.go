package clients

import (
	"context"
	"modulo/internal/models"

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
	query := `SELECT EXISTS(SELECT 1 FROM clientes WHERE id_cliente = $1)`
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
	query := `SELECT EXISTS(SELECT 1 FROM clientes WHERE correo = $1)`
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
	query := `SELECT EXISTS(SELECT 1 FROM clientes WHERE documento = $1)`
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

// Consultar informacion de clientes
func (r *RepositoryClients) QueryClientInformation() ([]models.Client, error) {
	ctx := context.Background()
	query := `SELECT id_cliente, nombre, documento, correo, telefono, fecha_ingreso, 
		contrasena, estado FROM clientes`
	// Trae todos los registros
	data, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var dataClients []models.Client
	// Recorre cada registro uno por uno
	for data.Next() {
		var client models.Client
		// Se agregan a memoria a client para luego pasarlo a dataclient
		if err := data.Scan(&client.Id, &client.Name, &client.Document,
			&client.Gmail, &client.Phone, &client.EntryDate,
			&client.Password, &client.State); err != nil {
			return nil, err
		}
		dataClients = append(dataClients, client)
	}
	return dataClients, nil
}

// Actualizacion informacion cliente con contraseña incluida
func (r *RepositoryClients) QueryUpdateClient(id_client int, name, document, gmail, phone, password, state string) error {
	ctx := context.Background()
	query := `UPDATE clientes SET nombre = $1 documento = $2 correo = $3, telefono = $4, contrasena = $5, estado = $6 
	WHERE id_cliente = $7`
	if _, err := r.db.Exec(ctx, query, name, document, gmail, phone, password, state, id_client); err != nil {
		return err
	}
	return nil
}

// Actualizacion informacion cliente sin contraseña incluida
func (r *RepositoryClients) QueryUpdateClientNoPassword(id_client int, name, document, gmail, phone, state string) error {
	ctx := context.Background()
	query := `UPDATE clientes SET nombre = $1 documento = $2 correo = $3, telefono = $4, esatdo = $5
	WHERE id_cliente = $6`
	if _, err := r.db.Exec(ctx, query, name, document, gmail, phone, state, id_client); err != nil {
		return err
	}
	return nil
}
