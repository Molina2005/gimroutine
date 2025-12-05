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
	// URLS respuestas http
	// {id} : va enrutado con el id de chi.URLParam(r, "id")
	r.Post("/users", handler.CreateUsers)
	r.Get("/users/{id}", repo.ConsultUserInformation)
	r.Put("/users/{id}", handler.UpdateUsersInformation)
	r.Delete("/users/{id}", handler.DeleteUsers)
	// Servidor escuchando en el puerto 8080
	http.ListenAndServe(":8080", r)
}
