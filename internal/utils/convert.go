package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Hashear contraseña
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Funcion apra comparar hash de base de datos con contraseña que el usuario ingresa
func ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

}
