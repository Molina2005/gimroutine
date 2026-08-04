package login

import (
	"context"
	"modulo/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepoLogin struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *RepoLogin {
	return &RepoLogin{db: db}
}

// Verificacion existencia de usuario por correo
// Uso: solo cuando se quieren registrar y ya existe el correo en sistema
func (r RepoLogin) QueryuserExistsXEmail(email string) (bool, error) {
	// contexto que exije *pgxpool.Pool para consultas sql
	ctx := context.Background()
	var existsEmail bool
	query := `SELECT EXISTS(SELECT 1 FROM usuarios WHERE correo=$1)`
	err := r.db.QueryRow(ctx, query, email).Scan(&existsEmail)
	// Mensaje de usuario ya en sistema
	if err != nil {
		return false, err
	}
	return existsEmail, nil
}

// Crear funcion validar existencia por correo
func (r *RepoLogin) QueryClientExistsByGmail(Gmail string) (bool, error) {
	ctx := context.Background()
	var ExistsGmail bool
	query := `SELECT EXISTS(SELECT 1 FROM clientes WHERE correo = $1)`
	err := r.db.QueryRow(ctx, query, Gmail).Scan(&ExistsGmail)
	if err != nil {
		return false, err
	}
	return ExistsGmail, nil
}

// Consulta para obtener login usuarios
func (r *RepoLogin) QueryLogin(email string) (*models.Login, error) {
	ctx := context.Background()
	var DataLogin models.Login
	query := `SELECT id_usuarios, correo, contrasena, rol FROM usuarios
		WHERE correo = $1`
	if err := r.db.QueryRow(ctx, query, email).Scan(
		&DataLogin.Id,
		&DataLogin.Email,
		&DataLogin.Password,
		&DataLogin.Role,
	); err != nil {
		return nil, err
	}
	return &DataLogin, nil
}

// Consulta para obtener login clientes
func (r *RepoLogin) QueryLoginClient(email string) (*models.Login, error) {
	ctx := context.Background()
	var DataLogin models.Login
	query := `SELECT id_cliente, correo, contrasena FROM clientes
		WHERE correo = $1`
	if err := r.db.QueryRow(ctx, query, email).Scan(
		&DataLogin.Id,
		&DataLogin.Email,
		&DataLogin.Password,
	); err != nil {
		return nil, err
	}
	return &DataLogin, nil
}
