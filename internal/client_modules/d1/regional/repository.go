package regional

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Guarda siempre la misma conexion a db
type RepositoryRegional struct {
	db *pgxpool.Pool
}

// Creacion nuevo repositroio para poder guardar la verdader conexion en la struct Repository
func NewRepository(db *pgxpool.Pool) *RepositoryRegional {
	return &RepositoryRegional{db: db}
}

// Funcion para validar existencia de regional por id
func (r RepositoryRegional) QueryRegionalExistsXId(id_regional int) (bool, error) {
	ctx := context.Background()
	var existsId bool
	query := `SELECT EXISTS(SELECT 1 FROM regionales WHERE id_regional=$1)`
	err := r.db.QueryRow(ctx, query, id_regional).Scan(&existsId)
	if err != nil {
		return false, err
	}
	return existsId, nil
}

// Funcion para validar existencia de regional por nombre
func (r RepositoryRegional) QueryRegionalExistsXName(name string) (bool, error) {
	ctx := context.Background()
	var existsName bool
	query := `SELECT EXISTS(SELECT 1 FROM regionales WHERE nombre=$1)`
	err := r.db.QueryRow(ctx, query, name).Scan(&existsName)
	if err != nil {
		return false, err
	}
	return existsName, nil
}

// Insercion de regionales
func (r *RepositoryRegional) QueryInsertRegional(id_client, id_user_state int, name string) error {
	ctx := context.Background()
	query := `INSERT INTO regionales (id_cliente, nombre, id_estado_usuario) 
				VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, id_client, name, id_user_state)
	if err != nil {
		return err
	}
	return nil
}

// Consultar Informacion regional
func (r *RepositoryRegional) QueryViewRegionalInfomation(id_regional int) (*DataRegional, error) {
	ctx := context.Background()
	// Guarda la informacion que se envia con Scan
	var DataRegional DataRegional
	query := `SELECT r.nombre, c.id_cliente, c.nombre AS cliente, eu.id_estado_usuario, eu.nombre AS estado
		FROM regionales r
		JOIN clientes c ON c.id_cliente = r.id_cliente
		JOIN estados_usuarios eu ON eu.id_estado_usuario = r.id_estado_usuario
		WHERE id_regional = $1`
	// Envia informacion con scan a la variable DataRegional
	if err := r.db.QueryRow(ctx, query, id_regional).Scan(
		&DataRegional.Name,
		&DataRegional.IdClient,
		&DataRegional.Client,
		&DataRegional.IdUserState,
		&DataRegional.Users_state,
	); err != nil {
		return nil, err
	}
	// retorna la informacion o un error
	return &DataRegional, nil
}

// Consulta informacion de todas las regionales existentes
func (r *RepositoryRegional) QueryRegionalInformation() ([]AllRegionales, error) {
	ctx := context.Background()
	query := `SELECT r.id_regional, r.nombre, c.nombre AS cliente, eu.nombre AS estado
		FROM regionales r
		JOIN clientes c ON c.id_cliente = r.id_cliente
		JOIN estados_usuarios eu ON eu.id_estado_usuario = r.id_estado_usuario`
	data, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var DataRegional []AllRegionales
	for data.Next() {
		var Modelsregional AllRegionales
		if err := data.Scan(
			&Modelsregional.Id,
			&Modelsregional.Name,
			&Modelsregional.Client,
			&Modelsregional.State); err != nil {
			return nil, err
		}
		DataRegional = append(DataRegional, Modelsregional)
	}
	return DataRegional, nil
}

// Funcion para actualizar datos de regional
func (r *RepositoryRegional) QueryRegionalUpdate(id_regional, id_client, id_state int, name string) error {
	ctx := context.Background()
	query := `UPDATE regionales SET nombre = $1, id_cliente = $2, id_estado_usuario = $3
	WHERE id_regional = $4`
	if _, err := r.db.Exec(ctx, query, name, id_client, id_state, id_regional); err != nil {
		return err
	}
	return nil
}

// Funcion eliminar regional
func (r *RepositoryRegional) QueryRegionalDelete(id_regional int) error {
	ctx := context.Background()
	query := `DELETE FROM regionales WHERE id_regional = $1`
	_, err := r.db.Exec(ctx, query, id_regional)
	if err != nil {
		return err
	}
	return nil
}

// Funcion buscador de regionales
func (r *RepositoryRegional) QueryRegionalSearch(search string) ([]RegionalSearch, error) {
	ctx := context.Background()
	query := `SELECT r.id_regional, r.nombre, c.nombre AS cliente, eu.nombre AS estado 
		FROM regionales r
		JOIN clientes c ON c.id_cliente = r.id_cliente 
		JOIN estados_usuarios eu ON eu.id_estado_usuario = c.id_estado_usuario
		WHERE r.nombre ILIKE '%' || $1 || '%'`
	data, err := r.db.Query(ctx, query, search)
	if err != nil {
		return nil, err
	}
	var DataRegional []RegionalSearch
	defer data.Close()
	for data.Next() {
		var dataModels RegionalSearch
		err := data.Scan(&dataModels.Id, &dataModels.Name, &dataModels.Client,
			&dataModels.State)
		if err != nil {
			return nil, err
		}
		DataRegional = append(DataRegional, dataModels)
	}
	return DataRegional, nil
}
