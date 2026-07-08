package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// USUARIOS

// Constular informacion usuarios
type User struct {
	Name     string
	Email    string
	Password string
	Role     string
}

// Login
type Login struct {
	Id       int
	Email    string
	Password string
	Role     string
}

// Struct informacion usuarios
type Users struct {
	Id        int
	Name      string
	Gmail     string
	Password  string
	EntryDate string
	Role      string
}

type SearchUsers struct {
	Name  string
	Email string
}

// CLIENTES

// Struct informacion clientes
type Client struct {
	Id        int
	Name      string
	Document  string
	Gmail     string
	Phone     string
	EnterDate time.Time
	Password  string
	State     string
}

// Struct manejo token JWT
type DataToken struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
	// campo para poder usar funciones de JWT
	jwt.RegisteredClaims
}

// TIPOS MANTENIMIENTOS

// Struct informacion tipo de mantenimiento
type TeamOfMaintenance struct {
	Id           int
	Name         string
	Description  string
	CreationDate time.Time
}

// EJERCICIOS

// Structs informacion ejercicios
type Maintenance struct {
	Id                  int
	IdTeamOfMaintenance int
	Name                string
	Description         string
	Img                 string
	CreationDate        time.Time
}
