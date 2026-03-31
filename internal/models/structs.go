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

// Struct informacion tipo de ejercicios
type TypeOfExercises struct {
	Id           int
	Name         string
	Description  string
	CreationDate time.Time
}

// Structs informacion ejercicios
type Exercises struct {
	Id               int
	IdTypeOfExercise int
	Name             string
	Description      string
	Img              string
	CreationDate     time.Time
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
