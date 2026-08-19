package plans

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Guarda siempre la misma conexion a db
type RepositoryPlans struct {
	db *pgxpool.Pool
}

// Creacion nuevo repositroio para poder guardar la verdader conexion en la struct Repository
func NewRepository(db *pgxpool.Pool) *RepositoryPlans {
	return &RepositoryPlans{db: db}
}

// Validacion existencia de planes
func (r *RepositoryPlans) QueryPlanExistOfName(name string) (bool, error) {
	ctx := context.Background()
	var existPlan bool
	query := `SELECT EXISTS(SELECT 1 FROM planes WHERE name = $1) `

	if err := r.db.QueryRow(ctx, query, name).Scan(&existPlan); err != nil {
		return false, nil
	}
	return existPlan, nil
}

// Insercion de planes
func (r *RepositoryPlans) QueryInsertPlans(name, description string, price, durationMonths, maxUser int, expirationDate time.Time) error {
	ctx := context.Background()
	query := `INSERT INTO planes (nombre, descripcion, precio, duracion_meses, max_usuarios, fecha_vencimiento) 
				VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(ctx, query, name, description, price, durationMonths, maxUser, expirationDate)
	if err != nil {
		return err
	}
	return nil
}
