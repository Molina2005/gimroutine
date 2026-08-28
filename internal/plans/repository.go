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

// Insercion de planes datos basicos del plan
func (r *RepositoryPlans) QueryInsertPlans(name, description string, userMax int) (int, error) {
	ctx := context.Background()
	query := `INSERT INTO planes (nombre, descripcion, max_usuarios) 
				VALUES ($1, $2, $3)
				RETURNING id_plan`
	var IdPlan int
	err := r.db.QueryRow(ctx, query, name, description, userMax).Scan(&IdPlan)
	if err != nil {
		return 0, err
	}
	return IdPlan, nil
}

// Funcion creacion de plan, datos basicos de los precios para manejar [mensual, anual, etc]
// Insercion de planes datos basicos del plan
func (r *RepositoryPlans) QueryInsertPrice(id_plan, Months, Price int) error {
	ctx := context.Background()
	query := `INSERT INTO precios (id_plan, meses, precio)
				VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, id_plan, Months, Price)
	if err != nil {
		return err
	}
	return nil
}

// Consultar para llamar todos los planes creados
func (r *RepositoryPlans) QueryAllPlans() ([]models.JoinPlanAndPrice, error) {
	ctx := context.Background()
	query := `SELECT pns.nombre, pns.descripcion, pns.max_usuarios, pr.meses, pr.precio, pns.fecha_creacion FROM planes AS pns 
		JOIN precios AS pr ON pr.id_plan = pns.id_plan`
	data, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var DataPlans []models.JoinPlanAndPrice
	for data.Next() {
		var plans models.JoinPlanAndPrice
		if err := data.Scan(&plans.Name, &plans.Description, &plans.UserMax, &plans.Months, &plans.Price, &plans.CretionDate); err != nil {
			return nil, err
		}
		DataPlans = append(DataPlans, plans)
	}
	return DataPlans, nil
}

// Funcion para unir tablas de planes y precios de planes segun meses
func (r *RepositoryPlans) QueryPlansIndex() ([]models.JoinPlanAndPrice, error) {
	ctx := context.Background()
	query := `SELECT pns.nombre, pns.descripcion, pns.max_usuarios, pr.meses, pr.precio FROM planes AS pns
		JOIN precios AS pr ON pr.id_plan = pns.id_plan WHERE pr.meses = 1`
	data, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	var dataPlansAndPrice []models.JoinPlanAndPrice
	for data.Next() {
		var plansAndPrice models.JoinPlanAndPrice
		if err := data.Scan(&plansAndPrice.Name, &plansAndPrice.Description, &plansAndPrice.UserMax, &plansAndPrice.Months, &plansAndPrice.Price); err != nil {
			return nil, err
		}
		dataPlansAndPrice = append(dataPlansAndPrice, plansAndPrice)
	}
	return dataPlansAndPrice, nil
}
