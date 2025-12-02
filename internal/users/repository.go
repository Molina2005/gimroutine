package users

import (
	"context"
	"errors"
	"modulo/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Guarda siempre la misma conexion a db
type Repository struct {
	db *pgxpool.Pool
}

// Creacion nuevo repositroio para poder guardar la verdader conexion en la struct Repository
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

// Verificacion existencia de usuario
func (r Repository) userExists(email string) (bool, error) {
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

// Insercion de usuarios
func (r *Repository) InsertUser(name, email string, age int, weight int16, height float64, password string) error {
	ctx := context.Background()
	// Funcion para verificar existencia de usuario
	existsEmail, err := r.userExists(email)
	if err != nil {
		return err
	}
	if existsEmail {
		return errors.New("usuario ya existe en el sistema")
	}
	// Si no existe lo crea
	if !existsEmail {
		query := `INSERT INTO usuarios (nombre, correo, edad, peso, altura, contrasena) 
				VALUES ($1, $2, $3, $4, $5, $6 )`
		_, err := r.db.Exec(ctx, query, name, email, age, weight, height, password)
		return err
	}
	return nil
}

// Consultar Informacion usuario
func (r *Repository) ViewUserInfomation(id_usuarios int, email string) (*models.User, error) {
	ctx := context.Background()
	existsEmail, err := r.userExists(email)
	if err != nil {
		return nil, err
	}
	// Error de no existencia de usuario
	if !existsEmail {
		return nil, errors.New("usuario no existe en el sistema")
	}
	// Guarda la informacion que se envia con Scan
	var DataUser models.User
	query := `SELECT id_usuarios,nombre,correo,edad,peso,altura,contrasena,fecha_ingreso
	FROM usuarios 
	WHERE id_usuarios = $1`
	// Envia informacion con scan a la variable DataUser
	if err := r.db.QueryRow(ctx, query, id_usuarios).Scan(
		&DataUser.Id_user,
		&DataUser.Name,
		&DataUser.Email,
		&DataUser.Age,
		&DataUser.Weight,
		&DataUser.Height,
		&DataUser.Password,
		&DataUser.EntryDate,
	); err != nil {
		return nil, err
	}
	// retorna la informacion o un error
	return &DataUser, nil
}

// Actualizacion informacion usuario
func (r Repository) UpdateUser(id_usuarios int, name, email string, age int, weight int16, height float64, password string) error {
	ctx := context.Background()
	query := `UPDATE usuarios SET nombre = $1, correo = $2, edad = $3, peso = $4, altura = $5, contrasena = $6
	WHERE id_usuarios = $7`
	_, err := r.db.Exec(ctx, query, name, email, age, weight, height, password, id_usuarios)
	if err != nil {
		return err
	}
	return nil
}

// Eliminar usuario del sistema
func (r Repository) DeleteUser() {

}
