package vitalincidents

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Guarda siempre la misma conexion a db
type RepositoryVitalIncidents struct {
	db *pgxpool.Pool
}

// Creacion nuevo repositroio para poder guardar la verdader conexion en la struct Repository
func NewRepository(db *pgxpool.Pool) *RepositoryVitalIncidents {
	return &RepositoryVitalIncidents{db: db}
}

// Consulta validar existencia de incidencia por id
func (r *RepositoryVitalIncidents) QueryVitalIncidentsCreate(numberIncidents int,
	servicePoint, equipment, description, technical, earrings, state, pdf, oc string,
	registrationDate, executionDate time.Time) error {

	return nil
}

// Consulta validar existencia de incidencia por numero de incidencia

// Consultacreacion de incidencias

// Consulta para obtener incidencia especifica

// Consulta obtener todas las incidencias segun la regional

// Consulta actualizacion informacion de incidencia

// Consulta eliminacion de incidencia

// Consulta buscador dinamico
