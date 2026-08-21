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

// Struct informacion usuarios
type Users struct {
	Id        int
	Name      string
	Gmail     string
	Password  string
	EntryDate string
	Role      string
}

// Struct search
type SearchUsers struct {
	Name      string
	Gmail     string
	Password  string
	EntryDate string
	Role      string
}

// CLIENTES

type Client struct {
	Name     string
	Document string
	Gmail    string
	Phone    string
	Password string
	State    string
}

// Struct informacion clientes
type Clients struct {
	Id        int
	Name      string
	Document  string
	Gmail     string
	Phone     string
	EntryDate time.Time
	Password  string
	State     string
}

// Struct search
type SearchClients struct {
	Name      string
	Document  string
	Gmail     string
	Phone     string
	Password  string
	EntryDate time.Time
	State     string
}

// LOGIN

// Login
type Login struct {
	Id       int
	Email    string
	Password string
	// Eleccion si es tipo [usuario] o [cliente]
	Type string
	Role string
}

// Struct manejo token JWT
type DataToken struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
	// campo para poder usar funciones de JWT
	jwt.RegisteredClaims
}

// PLANES

// maanejo de datos para llamar todos los campos de un plan
type PlansAll struct {
	Id          int
	Name        string
	Description string
	UserMax     int
	CretionDate time.Time
}
