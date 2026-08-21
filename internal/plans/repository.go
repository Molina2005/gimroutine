package plans

import (
	"context"
	"modulo/internal/models"

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
func (r *RepositoryPlans) QueryInsertPlans(name, description string, userMax int) error {
	ctx := context.Background()
	query := `INSERT INTO planes (nombre, descripcion, max_usuarios) 
				VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, name, description, userMax)
	if err != nil {
		return err
	}
	return nil
}

// Consultar para llamar todos los planes creados
func (r *RepositoryPlans) QueryAllPlans() ([]models.PlansAll, error) {
	ctx := context.Background()
	query := `SELECT id_plan, nombre, descripcion, max_usuarios, fecha_creacion
	FROM planes ORDER BY id_plan ASC`
	data, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var DataPlans []models.PlansAll
	for data.Next() {
		var plans models.PlansAll
		if err := data.Scan(&plans.Id, &plans.Name, &plans.Description, &plans.UserMax, &plans.CretionDate); err != nil {
			return nil, err
		}
		DataPlans = append(DataPlans, plans)
	}
	return DataPlans, nil
}
