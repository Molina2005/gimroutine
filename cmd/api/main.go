package main

import (
	"fmt"
	"log"
	teammaintenance "modulo/internal/TeamMaintenance"
	"modulo/internal/clients"
	"modulo/internal/database"
	maintenances "modulo/internal/maintenance"
	"modulo/internal/users"
	"net/http"
	"os"

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
	// LLamado de las capas para poder manipular equipos de mantenimiento
	repoTeamOfMaintenance := teammaintenance.NewRepository(connect)
	serviceTeamOfMaintenance := teammaintenance.NewService(repoTeamOfMaintenance)
	handlerTeamOfMaintenance := teammaintenance.NewHandler(serviceTeamOfMaintenance)
	// Llamado capas para poder manipular mantenimientos
	repoMaintenance := maintenances.NewRepository(connect)
	serviceMaintenance := maintenances.NewService(repoMaintenance)
	handlerMaintenance := maintenances.NewHanlder(serviceMaintenance)
	// Llamado capas para poder manipular clientes
	repoClients := clients.NewRepository(connect)
	serviceClients := clients.NewService(repoClients)
	handlerClients := clients.NewHandler(serviceClients)
	// Creacion de nuevo enrutador
	r := chi.NewRouter()
	// URLS respuestas http (usuario)
	// {id} : va enrutado con el id de chi.URLParam(r, "id")
	r.Post("/users", handler.HandlerCreateUsers)
	r.Post("/login", handler.HandlerLogin)
	r.Get("/users/{id}", handler.HandlerConsultUserInformation)
	r.Put("/users/{id}", handler.HandlerUpdateUsersInformation)
	r.Delete("/users/{id}", handler.HandlerDeleteUsers)
	r.Get("/allUsers", handler.HandlerConsultAllUsers)
	// URLS respuestas http (equipos de mantenimientos)
	r.Post("/TeamOfMaintenance", handlerTeamOfMaintenance.HandlerCreationTeamOfMaintenance)
	r.Get("/TeamOfMaintenance/{id}", handlerTeamOfMaintenance.HandlerConsultTeamOfMaintenance)
	r.Put("/TeamOfMaintenance/{id}", handlerTeamOfMaintenance.HandlerUpdateInfoTeamOfMaintenance)
	r.Delete("/TeaemOfMaintenance/{id}", handlerTeamOfMaintenance.HandlerDeleteTeamOfMaintenance)
	// URLS respuestas http (mantenimientos)
	r.Post("/Maintenance", handlerMaintenance.HandlerCreationMaintenance)
	r.Get("/Maintenance/{id}", handlerMaintenance.HandlerConsultInformationMaintenance)
	r.Put("/Maintenance/{id}", handlerMaintenance.HandlerUpdateInformationMaintenance)
	r.Delete("/Maintenance/{id}", handlerMaintenance.HandlerDeleteMaintenance)
	// URSL respuestas htpp (clientes)
	r.Post("/AddClients", handlerClients.HandlerCreationClients)
	r.Get("/AllClients", handlerClients.HandlerConsultAllClients)
	// Permite establecer la ruta html del index
	r.Get("/index", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/index.html")
	})
	// Permite establecer la ruta html de crear usuario
	r.Get("/createUsers", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/admin/forms/users/add_users.html")
	})
	// Permite establecer la ruta html del login
	r.Get("/loginUsers", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/users/login.html")
	})
	// Permite establecer la ruta html de pagina de informacion de usuarios
	r.Get("/usersInfo", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/users/users.html")
	})
	r.Get("/FormUpdateUsers", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/admin/forms/users/updateUsers.html")
	})
	// Pagina inicial para el admin
	r.Get("/adminHome", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/adminHome.html")
	})
	// Pagina inicial para el supervisor
	r.Get("/supervisorHome", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/supervisorHome.html")
	})
	// Pagina inicial para el tecnico
	r.Get("/TechnicalHome", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/technicalHome.html")
	})
	// pagina navegador general
	r.Get("/nav", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/components/navGeneral.html")
	})
	// pagina formulario politica y privacidad
	r.Get("/PrivacyPolicy", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/politica_privacidad.html")
	})
	// Pagina informacion clientes
	r.Get("/clientsInfo", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/admin/clients.html")
	})
	// Pagina agregar cliente
	r.Get("/addClient", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/admin/forms/clients/add_client.html")
	})
	// Ruta para poder trabajar con los archivos statics como css, js, etc...
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))
	// Servidor escuchando en el puerto XXXX
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, r)
}
