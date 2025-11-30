package users

import (
	"context"
	// "time"

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
	query := `INSERT INTO usuarios (nombre, correo, edad, peso, altura, contrasena) 
				VALUES ($1, $2, $3, $4, $5, $6 )`
	_, err := r.db.Exec(ctx, query, name, email, age, weight, height, password)
	return err
}
