package users

import (
	"context"
	"errors"

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

// Consulta insercion de usuarios
func (r *Repository) InsertUser(name, email string, age int, weight int16, height float64, password string) error {
	// contexto que exije *pgxpool.Pool para consultas sql
	ctx := context.Background()
	// Consulta de no creacion de usuario existente
	var existsEmail bool
	err := r.db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM usuarios WHERE correo=$1)", email).Scan(&existsEmail)
	if err != nil {
		return err
	}
	// Mensaje de usuario ya en sistema
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
