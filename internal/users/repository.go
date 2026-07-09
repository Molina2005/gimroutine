package users

import (
	"context"
	"fmt"
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
	query := `SELECT EXISTS(SELECT 1 FROM usuarios WHERE correo=$1)`
	err := r.db.QueryRow(ctx, query, email).Scan(&existsEmail)
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
	query := `SELECT EXISTS(SELECT 1 FROM usuarios WHERE id_usuarios=$1)`
	err := r.db.QueryRow(ctx, query, id_user).Scan(&existsId)
	if err != nil {
		return false, err
	}
	return existsId, nil
}

// Insercion de usuarios
func (r *RepositoryUsers) QueryInsertUser(name, email, password, role string) error {
	ctx := context.Background()
	query := `INSERT INTO usuarios (nombre, correo, contrasena, rol) 
				VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query, name, email, password, role)
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
	query := "SELECT nombre, correo, contrasena, rol FROM usuarios WHERE id_usuarios=$1"
	// Envia informacion con scan a la variable DataUser
	if err := r.db.QueryRow(ctx, query, id_user).Scan(
		&DataUser.Name,
		&DataUser.Email,
		&DataUser.Password,
		&DataUser.Role,
	); err != nil {
		return nil, err
	}
	// retorna la informacion o un error
	return &DataUser, nil
}

// Actualizacion informacion usuario con contraseña incluida
func (r *RepositoryUsers) QueryUpdateUser(id_user int, name, email, password, role string) error {
	ctx := context.Background()
	query := `UPDATE usuarios SET nombre = $1, correo = $2, contrasena = $3, rol = $4
	WHERE id_usuarios = $5`
	_, err := r.db.Exec(ctx, query, name, email, password, role, id_user)
	if err != nil {
		return err
	}
	return nil
}

// Actualizacion informacion usuario sin contraseña
func (r *RepositoryUsers) QueryUpdateUserNoPassword(id_user int, name, email, role string) error {
	ctx := context.Background()
	query := `UPDATE usuarios SET nombre = $1, correo = $2, rol = $3
	WHERE id_usuarios = $4`
	_, err := r.db.Exec(ctx, query, name, email, role, id_user)
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
	query := `SELECT id_usuarios, correo, contrasena, rol FROM usuarios
		WHERE correo = $1`
	if err := r.db.QueryRow(ctx, query, email).Scan(
		&DataLogin.Id,
		&DataLogin.Email,
		&DataLogin.Password,
		&DataLogin.Role,
	); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &DataLogin, nil
}

// Consulta informacion de todos los usuarios
func (r *RepositoryUsers) QueryUsersInformation() ([]models.Users, error) {
	ctx := context.Background()
	query := `SELECT id_usuarios, nombre, correo, contrasena, fecha_ingreso, rol FROM usuarios`
	data, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var DataUsers []models.Users
	for data.Next() {
		var users models.Users
		if err := data.Scan(&users.Id, &users.Name, &users.Gmail, &users.Password, &users.EntryDate, &users.Role); err != nil {
			return nil, err
		}
		DataUsers = append(DataUsers, users)
	}
	return DataUsers, nil
}

// Buscar usuario en buscador
func (r *RepositoryUsers) QuerySearchUsers(search string) ([]models.SearchUsers, error) {
	ctx := context.Background()
	query := `SELECT nombre, correo, contrasena, fecha_ingreso, rol 
	FROM usuarios 
	WHERE nombre ILIKE '%' || $1 || '%' 
	OR correo ILIKE '%' || $1 || '%'`
	data, err := r.db.Query(ctx, query, search)
	if err != nil {
		return nil, err
	}
	var DataUsers []models.SearchUsers
	defer data.Close()
	for data.Next() {
		var User models.SearchUsers
		err := data.Scan(&User.Name, &User.Email, &User.EntryDate, &User.Password, &User.Password)
		if err != nil {
			return nil, err
		}
		DataUsers = append(DataUsers, User)
	}
	return DataUsers, nil
}
