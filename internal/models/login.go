package models

import "github.com/golang-jwt/jwt/v5"

// Login
type Login struct {
	Id       int
	Email    string
	Password string
	// Eleccion si es tipo [usuario] o [cliente]
	Type   string
	IdRole int
}

// TOKEN JWT

// Struct manejo token JWT
type DataToken struct {
	Id     int `json:"id"`
	IdRole int `json:"idrole"`
	// campo para poder usar funciones de JWT
	jwt.RegisteredClaims
}
