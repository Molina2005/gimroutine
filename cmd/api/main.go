package main

import (
	"fmt"
	"log"
	"modulo/internal/config"
	"modulo/internal/database"
	"modulo/internal/handlers"
	httpmiddlewares "modulo/internal/http_middlewares"
	"modulo/internal/routes"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	// CONEXION A LA BASE DE DATOS
	connect, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer connect.Close()
	fmt.Println("CONEXION EXITOSA A LA BASE DE DATOS")
	// ENRUTADOR PARA MANEJAR INFORMACION
	r := chi.NewRouter()
	// LLAMADO DE CORS PARA PERMISOS DE DOMINIOS BACKEND Y FRONTEND
	r.Use(httpmiddlewares.EnableCors)
	// CAPAS SEGUN MODULOS
	config.LayersOfModules(connect)
	// ENRUTADORES PARA MANEJO DE HANDLER
	routes.Routes(connect, r)
	// LLAMADO DE FUNCION PARA CARGAR CADA UNA DE LAS PAGINAS ENRUTADORAS
	r.Mount("/", handlers.HandlerPages())
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))

	// SERVIDOR ESCUCHANDO EN PUERTO XXXX
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, r)
}
