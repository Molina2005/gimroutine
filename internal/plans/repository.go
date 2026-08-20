package plans

import (
	"context"

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
	query := `SELECT EXISTS(SELECT 1 FROM planes WHERE nombre = $1) `

	if err := r.db.QueryRow(ctx, query, name).Scan(&existPlan); err != nil {
		return false, err
	}
	return existPlan, nil
}

// Insercion de planes
func (r *RepositoryPlans) QueryInsertPlans(name, description string, price, monthsduration, userMax int) error {
	ctx := context.Background()
	query := `INSERT INTO planes (nombre, descripcion, precio, duracion_meses, max_usuarios) 
				VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query, name, description, price, monthsduration, userMax)
	if err != nil {
		return err
	}
	return nil
}
