package users

import (
	"database/sql"
	"time"
)

// Guarda siempre la misma conexion a db
type Repository struct {
	db *sql.DB
}

// Creacion nuevo repositroio para poder guardar la verdader conexion en la struct Repository
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// Consulta insercion de usuarios
func (r *Repository) InsertUser(name, email string, age int, weight int16, height float64, entryDate time.Time, password string) error {
	query := `INSERT INTO usuarios (nombre, correo, edad, peso, altura, fecha_ingreso, contrasena) 
				VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, name, email, age, weight, height, entryDate, password)
	return err
}
