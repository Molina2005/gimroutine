package config

import (
	jobposition "modulo/internal/client_modules/d1/jobPosition"
	"modulo/internal/client_modules/d1/regional"
	regionalcontacts "modulo/internal/client_modules/d1/regionalContacts"
	servicepoints "modulo/internal/client_modules/d1/servicePoints"
	"modulo/internal/clients"
	incidencestate "modulo/internal/incidence_state"
	"modulo/internal/login"
	"modulo/internal/roles"
	"modulo/internal/users"
	usersstate "modulo/internal/users_state"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Estructura para poder manejar handler de cada uno de los modulos
type Promise struct {
	ClientHandler           *clients.HandlerClients
	IncidenceHandler        *incidencestate.HandlerIncidenceState
	LoginHandler            *login.HandlerLogin
	RolesHandler            *roles.HandlerRoles
	UsersHandler            *users.HandlerUsers
	StatesHandler           *usersstate.HandlerUserState
	RegionalHandler         *regional.HandlerRegional
	JobPositionHandler      *jobposition.HandlerJobPosition
	RegionalContactsHandler *regionalcontacts.HandlerRegionalContacts
	ServicePointsHandler    *servicepoints.HandlerServicePoints
}

// Manejo de capas de cada uno de los modulos existentes
func LayersOfModules(connect *pgxpool.Pool) *Promise {
	// LOGIN
	repoLogin := login.NewRepository(connect)
	serviceLogin := login.NewService(repoLogin)
	handlerlogin := login.NewHandler(serviceLogin)

	// ROLES
	repoRoles := roles.NewRepository(connect)
	serviceRoles := roles.NewService(repoRoles)
	handlerRoles := roles.NewHandler(serviceRoles)

	// USUARIOS
	repo := users.NewRepository(connect)
	service := users.NewService(repo, repoLogin)
	handlerUser := users.NewHandler(service)

	// CLIENTES
	repoClients := clients.NewRepository(connect)
	serviceClients := clients.NewService(repoClients, repoLogin)
	handlerClients := clients.NewHandler(serviceClients)

	// ESTADOS USUARIOS
	repoUserStates := usersstate.NewRepository(connect)
	serviceUserStates := usersstate.NewService(repoUserStates)
	handlerUserStates := usersstate.NewHandler(serviceUserStates)

	// REGIONALES
	repoRegional := regional.NewRepository(connect)
	serviceRegional := regional.NewService(repoRegional)
	handlerRegional := regional.NewHandler(serviceRegional)

	// CARGOS
	repoJobPosition := jobposition.NewRepository(connect)
	serviceJobPosition := jobposition.NewService(repoJobPosition)
	handlerJobPosition := jobposition.NewHandler(serviceJobPosition)

	// CONTACTOS REGIONALES
	repoRegionalContacts := regionalcontacts.NewRepository(connect)
	serviceRegionalContacts := regionalcontacts.NewService(repoRegionalContacts)
	handlerRegionalContacts := regionalcontacts.NewHandler(serviceRegionalContacts)

	// PUNTOS SERVICIOS
	repoServicePoints := servicepoints.NewRepository(connect)
	serviceServicePoints := servicepoints.NewService(repoServicePoints)
	handlerServicePoints := servicepoints.NewHandler(serviceServicePoints)
	return &Promise{
		LoginHandler:            handlerlogin,
		RolesHandler:            handlerRoles,
		UsersHandler:            handlerUser,
		ClientHandler:           handlerClients,
		StatesHandler:           handlerUserStates,
		RegionalHandler:         handlerRegional,
		JobPositionHandler:      handlerJobPosition,
		RegionalContactsHandler: handlerRegionalContacts,
		ServicePointsHandler:    handlerServicePoints,
	}
}
