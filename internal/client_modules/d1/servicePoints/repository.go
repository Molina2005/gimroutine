package servicepoints

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Guarda siempre la misma conexion a db
type RepositoryServicePoints struct {
	db *pgxpool.Pool
}

// Creacion nuevo repositroio para poder guardar la verdader conexion en la struct Repository
func NewRepository(db *pgxpool.Pool) *RepositoryServicePoints {
	return &RepositoryServicePoints{db: db}
}

// Funcion validacion de existencia de punto servicio por id
func (r *RepositoryServicePoints) ExistsServicePointsOfId(id int) (bool, error) {
	ctx := context.Background()
	var ExistsId bool
	query := `SELECT EXISTS (SELECT 1 FROM puntos_servicios WHERE id_punto_servicio = $1)`
	err := r.db.QueryRow(ctx, query, id).Scan(&ExistsId)
	if err != nil {
		return false, err
	}
	return ExistsId, nil
}

// Funcion validacion de existencia de punto servicio por nombre
func (r *RepositoryServicePoints) ExistsServicePointsOfName(name string) (bool, error) {
	ctx := context.Background()
	var ExistsName bool
	query := `SELECT EXISTS (SELECT 1 FROM puntos_servicios WHERE nombre = $1)`
	err := r.db.QueryRow(ctx, query, name).Scan(&ExistsName)
	if err != nil {
		return false, err
	}
	return ExistsName, nil
}

// Funcion creacion de puntos de servicios
func (r *RepositoryServicePoints) QueryServicePointCreate(name, addres, city string, idRegional, idState int) error {
	ctx := context.Background()
	query := `INSERT INTO puntos_servicios (nombre, direccion, ciudad, id_regional, id_estado_usuario) 
					VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query, name, addres, city, idRegional, idState)
	if err != nil {
		return err
	}
	return nil
}

// Funcion obtener todos los puntos de servicios
func (r *RepositoryServicePoints) QueryAllServicePoints(idRegional int) ([]ServicePoints, error) {
	ctx := context.Background()
	query := `SELECT ps.id_punto_servicio, ps.nombre AS nombre_punto, ps.direccion, ps.ciudad, eu.nombre AS estado FROM puntos_servicios ps 
		JOIN estados_usuarios eu ON eu.id_estado_usuario = ps.id_estado_usuario WHERE id_regional = $1`
	data, err := r.db.Query(ctx, query, idRegional)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var DataServicePoints []ServicePoints
	for data.Next() {
		var ModelsServicePoints ServicePoints
		if err := data.Scan(&ModelsServicePoints.Id, &ModelsServicePoints.Name, &ModelsServicePoints.Addres,
			&ModelsServicePoints.City, &ModelsServicePoints.State); err != nil {
			return nil, err
		}
		DataServicePoints = append(DataServicePoints, ModelsServicePoints)
	}
	return DataServicePoints, nil
}

// Funcion consultar y/o obtener punto servicio especifico
func (r *RepositoryServicePoints) QueryServicePointInformation(idServicePoint int) (*ServicePoint, error) {
	ctx := context.Background()
	var dataServicePoints ServicePoint
	query := `SELECT ps.nombre AS nombre_punto, ps.direccion, ps.ciudad, ps.id_estado_usuario FROM puntos_servicios ps
	JOIN estados_usuarios eu ON eu.id_estado_usuario = ps.id_estado_usuario	WHERE ps.id_punto_servicio = $1`
	if err := r.db.QueryRow(ctx, query, idServicePoint).Scan(
		&dataServicePoints.Name, &dataServicePoints.Addres,
		&dataServicePoints.City, &dataServicePoints.IdState); err != nil {
		return nil, err
	}
	return &dataServicePoints, nil
}

// Funcion actualizacion de informacion de punto servicio
