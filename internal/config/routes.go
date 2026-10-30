package config

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Funcion para manejar handlers de cada uno de susu modulos, con su URL de acceso
func Routes(connect *pgxpool.Pool, r *chi.Mux) {
	// Llamado funcion para obtener informacion de handlers de modulos
	promises := LayersOfModules(connect)

	// ROLES
	r.Get("/Allroles", promises.RolesHandler.HandlerConsultAllRoles)

	// USUARIOS
	r.Post("/users", promises.UsersHandler.HandlerCreateUsers)
	r.Get("/users/{id}", promises.UsersHandler.HandlerConsultUserInformation)
	r.Put("/users/{id}", promises.UsersHandler.HandlerUpdateUsersInformation)
	r.Delete("/users/{id}", promises.UsersHandler.HandlerDeleteUsers)
	r.Get("/allUsers", promises.UsersHandler.HandlerConsultAllUsers)
	r.Get("/UsersSearch/{search}", promises.UsersHandler.HandlerUsersSearch)

	// CLIENTES
	r.Post("/AddClients", promises.ClientHandler.HandlerCreationClients)
	r.Get("/client/{id}", promises.ClientHandler.HandlerConsultClientInformation)
	r.Get("/AllClients", promises.ClientHandler.HandlerConsultAllClients)
	r.Put("/client/{id}", promises.ClientHandler.HandlerUpdateClients)
	r.Delete("/client/{id}", promises.ClientHandler.HandlerDeleteClient)
	r.Get("/ClientSearch/{search}", promises.ClientHandler.HandlerClientsSearch)
	r.Get("/clientPage/{id}", promises.ClientHandler.HandlerClientsPages)

	// ESTADOS
	r.Get("/AllUserStates", promises.StatesHandler.HandlerAllUserState)

	// LOGIN
	r.Post("/login", promises.LoginHandler.HandlerLogin)

	// REGIONALES
	r.Post("/regional", promises.RegionalHandler.HandlerCreateRegional)
	r.Get("/regional/{id}", promises.RegionalHandler.HandlerConsultUserInformation)
	r.Get("/allRegional", promises.RegionalHandler.HandlerConsultAllRegional)
	r.Put("/regional/{id}", promises.RegionalHandler.HandlerUpdateRegional)
	r.Delete("/regional/{id}", promises.RegionalHandler.HandlerDeleteRegional)
	r.Get("/regionalSearch/{search}", promises.RegionalHandler.HandlerRegionalSearch)

	// CARGO DE TRABAJO
	r.Get("/AllJobPosition", promises.JobPositionHandler.HandlerAllJobPosition)

	// CONTACTOS REGIONALES
	r.Post("/regionalContacts", promises.RegionalContactsHandler.HandlerRegionalContactsCreate)
	r.Get("/regionalContact/{id}", promises.RegionalContactsHandler.HandlerRegionalContacts)
	r.Get("/AllregionalContacts/{idRegional}", promises.RegionalContactsHandler.HandlerConsultAllContactsRegional)
	r.Put("/regionalContacts/{id}", promises.RegionalContactsHandler.HandlerRegionalContactsUpdate)
	r.Delete("/regionalContacts/{id}", promises.RegionalContactsHandler.HandlerRegionalContactsDelete)
	r.Get("/regionalContacts/{search}/{idRegional}", promises.RegionalContactsHandler.HandlerRegionalContactsSearch)

	// PUNTOS SERVICIOS
	r.Post("/servicePoints", promises.ServicePointsHandler.HandlerServicePointsCreate)
	r.Get("/servicePoints/{idRegional}", promises.ServicePointsHandler.HandlerAllServicePoints)
	r.Get("/servicePoint/{id}", promises.ServicePointsHandler.HandlerServicePointInformation)
	r.Put("/servicePoint/{id}", promises.ServicePointsHandler.HandlerServicePointsUpdate)
	r.Get("/servicePoints/{search}/{idRegional}", promises.ServicePointsHandler.HandlerServicePointSearch)
	r.Delete("/servicePoint/{id}", promises.ServicePointsHandler.HandlerServicePointsDelete)
}
