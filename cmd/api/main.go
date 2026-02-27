package main

import (
	"fmt"
	"log"
	typeexercises "modulo/internal/TypeExercises"
	"modulo/internal/database"
	"modulo/internal/exercises"
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
	// Llamado capas para poder manipular ejercicios
	repoExercises := exercises.NewRepository(connect)
	serviceExercises := exercises.NewService(repoExercises)
	handlerExercises := exercises.NewHanlder(serviceExercises)
	// Creacion de nuevo enrutador
	r := chi.NewRouter()
	// URLS respuestas http (usuario)
	// {id} : va enrutado con el id de chi.URLParam(r, "id")
	r.Post("/users", handler.HandlerCreateUsers)
	r.Post("/login", handler.HandlerLogin)
	r.Get("/users/{id}", repo.HandlerConsultUserInformation)
	r.Put("/users/{id}", handler.HandlerUpdateUsersInformation)
	r.Delete("/users/{id}", handler.HandlerDeleteUsers)
	// URLS respuestas http (tipos de ejercicios)
	r.Post("/TypeOfExercises", handlerTypeExercises.HandlerCreationTypeOfExercise)
	r.Get("/TypeOfExercises/{id}", repoTypeExercises.HandlerConsultTypeOfExercises)
	r.Put("/TypeOfExercises/{id}", handlerTypeExercises.HandlerUpdateInfoTypeOfExercises)
	r.Delete("/TypeOfExercises/{id}", handlerTypeExercises.HandlerDeleteTypeOfExercises)
	// URLS respuestas http (ejercicios)
	r.Post("/Exercises", handlerExercises.HandlerCreationExercises)
	r.Get("/Exercises/{id}", repoExercises.HandlerConsultInformationExercise)
	r.Put("/Exercises/{id}", handlerExercises.HandlerUpdateInformationExercise)
	r.Delete("/Exercises/{id}", handlerExercises.HandlerDeleteExercise)
	// Permite establecer la ruta html del index
	r.Get("/index", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/index.html")
	})
	// Permite establecer la ruta html de crear usuario
	r.Get("/createAccount", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/users/create_account.html")
	})
	// Permite establecer la ruta html del login
	r.Get("/loginUsers", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/users/login.html")
	})
	// Pagina inicial al iniciar sesion
	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/home.html")
	})
	r.Get("/nav", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/users/navGeneral.html")
	})
	// Ruta para poder trabajar con los archivos statics como css, js, etc...
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))
	// Servidor escuchando en el puerto 8080
	http.ListenAndServe(":3800", r)
}
