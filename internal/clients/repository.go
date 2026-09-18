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

// Crear fncion valdar exixtencia por correo
func (r *RepositoryClients) QueryClientExistsByDocument(gmail string) (bool, error) {
	ctx := context.Background()
	var ExistsGmail bool
	query := `SELECT EXISTS(SELECT 1 FROM clientes WHERE correo = $1)`
	err := r.db.QueryRow(ctx, query, gmail).Scan(&ExistsGmail)
	if err != nil {
		return false, err
	}
	return ExistsGmail, nil
}

// consulta para creacion de clientes
func (r *RepositoryClients) QueryCreateClient(name, phone, gmail, nit string, id_state int) error {
	ctx := context.Background()
	query := `INSERT INTO clientes (nombre,telefono,nit,correo, id_estado_usuario) 
		VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query, name, phone, nit, gmail, id_state)
	if err != nil {
		return err
	}
	return nil
}

// Consultar Informacion cliente
func (r *RepositoryClients) QueryViewClientInfomation(id_client int) (*models.Client, error) {
	ctx := context.Background()
	// Guarda la informacion que se envia con Scan
	var DataClient models.Client
	query := `SELECT c.nombre, c.telefono, eu.nombre AS estado, c.nit,
		correo FROM clientes c
		JOIN estados_usuarios eu ON eu.id_estado_usuario = c.id_estado_usuario 
		WHERE id_cliente=$1`
	// Envia informacion con scan a la variable DataClient
	if err := r.db.QueryRow(ctx, query, id_client).Scan(
		&DataClient.Name,
		&DataClient.Phone,
		&DataClient.State,
		&DataClient.Nit,
		&DataClient.Gmail,
	); err != nil {
		return nil, err
	}
	// retorna la informacion o un error
	return &DataClient, nil
}

// Consultar informacion de clientes
func (r *RepositoryClients) QueryClientInformation() ([]models.Clients, error) {
	ctx := context.Background()
	query := `SELECT c.id_cliente, c.nombre, c.telefono, c.fecha_ingreso, eu.nombre AS estado, c.nit,
	correo FROM clientes c
	JOIN estados_usuarios eu ON eu.id_estado_usuario = c.id_estado_usuario
	ORDER BY id_cliente ASC`
	// Trae todos los registros
	data, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var dataClients []models.Clients
	// Recorre cada registro uno por uno
	for data.Next() {
		var client models.Clients
		// Se agregan a memoria a client para luego pasarlo a dataclient
		if err := data.Scan(&client.Id, &client.Name, &client.Phone, &client.EntryDate, &client.State,
			&client.Nit, &client.Gmail); err != nil {
			return nil, err
		}
		dataClients = append(dataClients, client)
	}
	return dataClients, nil
}

// Actualizacion informacion cliente
func (r *RepositoryClients) QueryUpdateClient(id_client int, name, gmail, phone, nit string, id_state int) error {
	ctx := context.Background()
	query := `UPDATE clientes SET nombre = $1, correo = $2, telefono = $3, id_estado_usuario = $4, nit = $5
	WHERE id_cliente = $6`
	if _, err := r.db.Exec(ctx, query, name, gmail, phone, id_state, nit, id_client); err != nil {
		return err
	}
	return nil
}

// Eliminar cliente del sistema
func (r *RepositoryClients) QueryDeleteClient(id_client int) error {
	ctx := context.Background()
	query := `DELETE FROM clientes WHERE id_cliente = $1`
	_, err := r.db.Exec(ctx, query, id_client)
	if err != nil {
		return err
	}
	return nil
}

// Buscar cliente en buscador
func (r *RepositoryClients) QuerySearchClients(search string) ([]models.SearchClients, error) {
	ctx := context.Background()
	query := `SELECT c.nombre, c.telefono, c.fecha_ingreso, eu.nombre AS estado, c.nit,
	correo FROM clientes c
	JOIN estados_usuarios eu ON eu.id_estado_usuario = c.id_estado_usuario
	WHERE c.nombre ILIKE '%' || $1 || '%' 
	OR c.nit ILIKE '%' || $1 || '%' 
	OR c.correo ILIKE '%' || $1 || '%'`
	data, err := r.db.Query(ctx, query, search)
	if err != nil {
		return nil, err
	}
	var DataClients []models.SearchClients
	defer data.Close()
	for data.Next() {
		var User models.SearchClients
		err := data.Scan(&User.Name, &User.Phone, &User.EntryDate, &User.State, &User.Nit, &User.Gmail)
		if err != nil {
			return nil, err
		}
		DataClients = append(DataClients, User)
	}
	return DataClients, nil
}
