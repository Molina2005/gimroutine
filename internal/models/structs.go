package models

import "time"

type User struct {
	Id_user   int
	Name      string
	Email     string
	Age       int
	Weight    int16
	Height    float64
	Password  string
	EntryDate time.Time
}
