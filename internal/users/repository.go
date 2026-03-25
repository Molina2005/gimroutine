package users

import (
	"context"
	"modulo/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Guarda siempre la misma conexion a db
type RepositoryUsers struct {
	db *pgxpool.Pool
}

// Creacion nuevo repositroio para poder guardar la verdader conexion en la struct Repository
func NewRepository(db *pgxpool.Pool) *RepositoryUsers {
	return &RepositoryUsers{db: db}
}

// Verificacion existencia de usuario por correo
// Uso: solo cuando se quieren registrar y ya existe el correo en sistema
func (r RepositoryUsers) QueryuserExistsXEmail(email string) (bool, error) {
	// contexto que exije *pgxpool.Pool para consultas sql
	ctx := context.Background()
	var existsEmail bool
	err := r.db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM usuarios WHERE correo=$1)", email).Scan(&existsEmail)
	// Mensaje de usuario ya en sistema
	if err != nil {
		return false, err
	}
	return existsEmail, nil
}

// Verificacion existencia de usuario por id
// Uso: cuando ya estan en el sistema y se quiere validar que en realidad si esten por id
func (r RepositoryUsers) QueryuserExistsXId(id_user int) (bool, error) {
	ctx := context.Background()
	var existsId bool
	err := r.db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM usuarios WHERE id_usuarios=$1)", id_user).Scan(&existsId)
	if err != nil {
		return false, err
	}
	return existsId, nil
}

// Insercion de usuarios
func (r *RepositoryUsers) QueryInsertUser(name, email string, password string) error {
	ctx := context.Background()
	query := `INSERT INTO usuarios (nombre, correo, contrasena) 
				VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, name, email, password)
	if err != nil {
		return err
	}
	return nil
}

// Consultar Informacion usuario
func (r *RepositoryUsers) QueryViewUserInfomation(id_user int) (*models.User, error) {
	ctx := context.Background()
	// Guarda la informacion que se envia con Scan
	var DataUser models.User
	query := `SELECT * FROM usuarios WHERE id_usuarios = $1`
	// Envia informacion con scan a la variable DataUser
	if err := r.db.QueryRow(ctx, query, id_user).Scan(
		&DataUser.Id,
		&DataUser.Name,
		&DataUser.Email,
		&DataUser.Password,
		&DataUser.EntryDate,
	); err != nil {
		return nil, err
	}
	// retorna la informacion o un error
	return &DataUser, nil
}

// Actualizacion informacion usuario
func (r *RepositoryUsers) QueryUpdateUser(id_user int, name, email string, password string) error {
	ctx := context.Background()
	query := `UPDATE usuarios SET nombre = $1, correo = $2, contrasena = $3
	WHERE id_usuarios = $4`
	_, err := r.db.Exec(ctx, query, name, email, password, id_user)
	if err != nil {
		return err
	}
	return nil
}

// Eliminar usuario del sistema
func (r *RepositoryUsers) QueryDeleteUser(id_user int) error {
	ctx := context.Background()
	query := `DELETE FROM usuarios WHERE id_usuarios = $1`
	_, err := r.db.Exec(ctx, query, id_user)
	if err != nil {
		return err
	}
	return nil
}

// Consulta para obtener login
func (r *RepositoryUsers) QueryLogin(email string) (*models.Login, error) {
	ctx := context.Background()
	var DataLogin models.Login
	query := `SELECT correo, contrasena FROM usuarios
		WHERE correo = $1`
	if err := r.db.QueryRow(ctx, query, email).Scan(
		&DataLogin.Email,
		&DataLogin.Password,
	); err != nil {
		return nil, err
	}
	return &DataLogin, nil
}
