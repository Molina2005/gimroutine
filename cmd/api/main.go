package main

import (
	"fmt"
	"log"
	"modulo/internal/database"
	"modulo/internal/users"
	"net/http"
)

func main() {
	// Conexion a la base de datos
	connect, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer connect.Close()
	fmt.Println("conexion exitosa a postgres")

	// LLamado de las capas de users para poder crear usuarios
	repo := users.NewRepository(connect)
	service := users.NewService(repo)
	handler := users.NewHandler(service)
	http.HandleFunc("/users", handler.CreateUsers)
	http.ListenAndServe(":8080", nil)
}
