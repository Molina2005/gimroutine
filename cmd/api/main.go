package main

import (
	"fmt"
	"log"
	jobposition "modulo/internal/client_modules/d1/jobPosition"
	"modulo/internal/client_modules/d1/regional"
	regionalcontacts "modulo/internal/client_modules/d1/regionalContacts"
	"modulo/internal/clients"
	"modulo/internal/database"
	"modulo/internal/login"
	"modulo/internal/roles"
	"modulo/internal/users"
	usersstate "modulo/internal/users_state"
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
	// Llamado capas para poder manipular roles
	repoRoles := roles.NewRepository(connect)
	serviceRoles := roles.NewService(repoRoles)
	handlerRoles := roles.NewHandler(serviceRoles)
	// LLamado de las capas para poder maninpular usuarios
	repo := users.NewRepository(connect)
	service := users.NewService(repo, repoLogin)
	handler := users.NewHandler(service)
	// Llamado capas para poder manipular clientes
	repoClients := clients.NewRepository(connect)
	serviceClients := clients.NewService(repoClients, repoLogin)
	handlerClients := clients.NewHandler(serviceClients)
	// Llamado capas para poder manipular estados de usuarios
	repoUserStates := usersstate.NewRepository(connect)
	serviceUserStates := usersstate.NewService(repoUserStates)
	handlerUserStates := usersstate.NewHandler(serviceUserStates)
	// Llamado capas para poder manipular regionales
	repoRegional := regional.NewRepository(connect)
	serviceRegional := regional.NewService(repoRegional)
	handlerRegional := regional.NewHandler(serviceRegional)
	// Llamado capas para poder manipular cargos de trabajo
	repoJobPosition := jobposition.NewRepository(connect)
	serviceJobPosition := jobposition.NewService(repoJobPosition)
	handlerJobPosition := jobposition.NewHandler(serviceJobPosition)
	// Llamado capas para poder manipular contactos regionales
	repoRegionalContacts := regionalcontacts.NewRepository(connect)
	serviceRegionalContacts := regionalcontacts.NewService(repoRegionalContacts)
	handlerRegionalContacts := regionalcontacts.NewHandler(serviceRegionalContacts)
	// Creacion de nuevo enrutador
	r := chi.NewRouter()
	// Creacion de superAdmin
	service.ServiceCreateSuperAdmin()
	// URLS respuestas http (roles)
	r.Get("/Allroles", handlerRoles.HandlerConsultAllRoles)
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
	r.Get("/clientPage/{id}", handlerClients.HandlerClientsPages)
	// URSL respuestas htpp (estados_usaurios)
	r.Get("/AllUserStates", handlerUserStates.HandlerAllUserState)
	// URLS respuestas http (Login)
	r.Post("/login", handlerlogin.HandlerLogin)
	// URLS respuestas http (regionales)
	r.Post("/regional", handlerRegional.HandlerCreateRegional)
	r.Get("/regional/{id}", handlerRegional.HandlerConsultUserInformation)
	r.Get("/allRegional", handlerRegional.HandlerConsultAllRegional)
	r.Put("/regional/{id}", handlerRegional.HandlerUpdateRegional)
	r.Delete("/regional/{id}", handlerRegional.HandlerDeleteRegional)
	r.Get("/regionalSearch/{search}", handlerRegional.HandlerRegionalSearch)
	// URLS respuestas http (cargo de trabajo)
	r.Get("/AllJobPosition", handlerJobPosition.HandlerAllJobPosition)
	// URLS respuestas http (contactos regionales)
	r.Post("/regionalContacts", handlerRegionalContacts.HandlerRegionalContactsCreate)
	r.Get("/regionalContact/{id}", handlerRegionalContacts.HandlerRegionalContacts)
	r.Get("/AllregionalContacts/{idRegional}", handlerRegionalContacts.HandlerConsultAllContactsRegional)
	r.Put("/regionalContacts/{id}", handlerRegionalContacts.HandlerRegionalContactsUpdate)
	r.Delete("/regionalContacts/{id}", handlerRegionalContacts.HandlerRegionalContactsDelete)
	r.Get("/regionalContacts/{search}/{idRegional}", handlerRegionalContacts.HandlerRegionalContactsSearch)
	// Permite establecer la ruta html del index
	r.Get("/index", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/index.html")
	})
	// Permite establecer la ruta html de crear usuario
	r.Get("/createUsers", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/admin/forms/users/add_users.html")
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
	// Pagina inicial para el admin
	r.Get("/adminHome", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/homePages/adminHome.html")
	})
	// Pagina inicial para el clientes
	r.Get("/clientHome", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/homePages/clientHome.html")
	})
	// pagina navegador general
	r.Get("/nav", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/components/navGeneral.html")
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
	// Pagina listado de regionales existentes
	r.Get("/listRegional", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/lists/regional/regionalList.html")
	})
	// Pagina para crear regionales
	r.Get("/formCreateRegional", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/forms/regional/regionalCreate.html")
	})
	// Pagina para actualizar informacion regionales
	r.Get("/formUpdateRegional", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/forms/regional/regionalUpdate.html")
	})
	// pagina con menu de modulos para que usuario elija datos que desea ver segun la region
	r.Get("/moduleMenu", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/pages/moduleMenu.html")
	})
	// pagina para moder vizualizar los contactos de cada regional
	r.Get("/regionalContacts", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/lists/regionalContacts/regionalContacts.html")
	})
	// Pagina para creacion de nuevos contactos regional
	r.Get("/formCreateRegionalContacts", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/forms/regionalContacts/creationForm.html")
	})
	r.Get("/formUpdateRegionalContacts", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/forms/regionalContacts/updateForm.html")
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
