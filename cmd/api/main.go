package main

import (
	"fmt"
	"log"
	"modulo/internal/clients"
	"modulo/internal/database"
	"modulo/internal/login"
	"modulo/internal/plans"
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
	// Llamado capas para poder manipular LOGIN
	repoLogin := login.NewRepository(connect)
	serviceLogin := login.NewService(repoLogin)
	handlerlogin := login.NewHandler(serviceLogin)
	// LLamado de las capas para poder maninpular usuarios
	repo := users.NewRepository(connect)
	service := users.NewService(repo, repoLogin)
	handler := users.NewHandler(service)
	// Llamado capas para poder manipular clientes
	repoClients := clients.NewRepository(connect)
	serviceClients := clients.NewService(repoClients, repoLogin)
	handlerClients := clients.NewHandler(serviceClients)
	// Llamado capas para poder manipular planes
	repoPlans := plans.NewRepository(connect)
	servicePlans := plans.NewService(repoPlans)
	handlerPlans := plans.NewHandler(servicePlans)
	// Creacion de nuevo enrutador
	r := chi.NewRouter()
	// Creacion de superAdmin
	service.ServiceCreateSuperAdmin()

	// URLS respuestas http (usuario)
	// {id} : va enrutado con el id de chi.URLParam(r, "id")
	r.Post("/users", handler.HandlerCreateUsers)
	r.Get("/users/{id}", handler.HandlerConsultUserInformation)
	r.Put("/users/{id}", handler.HandlerUpdateUsersInformation)
	r.Delete("/users/{id}", handler.HandlerDeleteUsers)
	r.Get("/allUsers", handler.HandlerConsultAllUsers)
	r.Get("/UsersSearch/{search}", handler.HandlerUsersSearch)
	// URSL respuestas htpp (clientes)
	r.Post("/AddClients", handlerClients.HandlerCreationClients)
	r.Get("/client/{id}", handlerClients.HandlerConsultClientInformation)
	r.Get("/AllClients", handlerClients.HandlerConsultAllClients)
	r.Put("/client/{id}", handlerClients.HandlerUpdateClients)
	r.Delete("/client/{id}", handlerClients.HandlerDeleteClient)
	r.Get("/ClientSearch/{search}", handlerClients.HandlerClientsSearch)
	// URLS respuestas http (planes)
	r.Post("/addPlans", handlerPlans.HandlerCreatePlans)
	// URLS respuestas http (Login)
	r.Post("/login", handlerlogin.HandlerLogin)
	// Permite establecer la ruta html del index
	r.Get("/index", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/index.html")
	})
	// Permite establecer la ruta html de crear usuario
	r.Get("/createUsers", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/admin/forms/users/add_users.html")
	})
	// Permite establecer la ruta html de crear empresa
	r.Get("/CreateCompany", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/users/company.html")
	})
	// Permite establecer la ruta html del login
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/login/login.html")
	})
	// Permite establecer la ruta html de pagina de informacion de usuarios
	r.Get("/usersInfo", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/users/users.html")
	})
	// Permite ingresar al formulario de actualizacion de datos usuario
	r.Get("/FormUpdateUsers", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/admin/forms/users/updateUsers.html")
	})
	// Pagina inicial para el super admin
	r.Get("/superAdminHome", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/superAdmin.html")
	})
	// Pagina inicial para el admin
	r.Get("/adminHome", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/adminHome.html")
	})
	// Pagina inicial para el usaurios
	r.Get("/userHome", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/userHome.html")
	})
	// Pagina inicial para el clientes
	r.Get("/clientHome", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clientHome.html")
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
	// Pagina para actualizar informacion cliente
	r.Get("/FormUpdateClients", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/admin/forms/clients/updateClient.html")
	})

	// PENDIENTE INFORMACION QUE SE ESTABA REALIZANDO SOBRE LSO PLANES

	// pagina con lista de planes creados
	r.Get("/planList", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/plans/planList.html")
	})
	// pagina para crear planes
	r.Get("/planCreate", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/plans/forms/creationPlans.html")
	})
	// pagina actualizar informacion de planes
	r.Get("/planUpdate", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/plans/forms/updatePlans.html")
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
