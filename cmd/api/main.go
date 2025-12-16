package main

import (
	"fmt"
	"log"
	typeexercises "modulo/internal/TypeExercises"
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

	// LLamado de las capas para poder maninpular usuarios
	repo := users.NewRepository(connect)
	service := users.NewService(repo)
	handler := users.NewHandler(service)
	// LLamado de las capas para poder manipular tipos de ejercicios
	repoTypeExercises := typeexercises.NewRepository(connect)
	serviceTypeExercises := typeexercises.NewService(repoTypeExercises)
	handlerTypeExercises := typeexercises.NewHandler(serviceTypeExercises)

	// Creacion de nuevo enrutador
	r := chi.NewRouter()
	// URLS respuestas http (usuario)
	// {id} : va enrutado con el id de chi.URLParam(r, "id")
	r.Post("/users", handler.CreateUsers)
	r.Get("/users/{id}", repo.ConsultUserInformation)
	r.Put("/users/{id}", handler.UpdateUsersInformation)
	r.Delete("/users/{id}", handler.DeleteUsers)
	// URLS respuestas http (tipos de ejercicios)
	r.Post("/TypeOfExercises", handlerTypeExercises.HandlerCreationTypeOfExercise)
	r.Get("/TypeOfExercises/{id}", repoTypeExercises.HandlerConsultTypeOfExercises)
	r.Put("/TypeOfExercises/{id}", handlerTypeExercises.HandlerUpdateInfoTypeOfExercises)
	r.Delete("/TypeOfExercises/{id}", handlerTypeExercises.DeleteTypeOfExercises)
	// Servidor escuchando en el puerto 8080
	http.ListenAndServe(":8080", r)
}
