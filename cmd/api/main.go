package main

import (
	"fmt"
	"log"
	"modulo/internal/database"
	"modulo/internal/users"
	"net/http"

	"github.com/go-chi/chi/v5"
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

	// Creacion de nuevo enrutador
	r := chi.NewRouter()
	// URL creacion de usuarios
	r.Post("/users", handler.CreateUsers)
	// URL consultar informaicion usuarios
	// {id} : va enrutado con el id de chi.URLParam(r, "id")
	r.Get("/users/{id}", repo.ConsultUserInformation)
	// URL actualizar informacion usuario
	r.Put("/users/{id}", handler.UpdateUsersInformation)
	// Servidor escuchando en el puerto 8080
	http.ListenAndServe(":8080", r)
}
