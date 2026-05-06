package models

import "time"

// Struct Informacion usuarios
type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	EntryDate time.Time
}

// Login
type Login struct {
	Email    string
	Password string
}

// Struct informacion tipo de mantenimiento
type TeamOfMaintenance struct {
	Id           int
	Name         string
	Description  string
	CreationDate time.Time
}

// Structs informacion ejercicios
type Maintenance struct {
	Id                  int
	IdTeamOfMaintenance int
	Name                string
	Description         string
	Img                 string
	CreationDate        time.Time
}

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
