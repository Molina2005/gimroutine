package models

import "time"

// Struct Informacion usuarios
type User struct {
	Id        int
	Name      string
	Email     string
	Age       int
	Weight    int16
	Height    float64
	Password  string
	EntryDate time.Time
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
