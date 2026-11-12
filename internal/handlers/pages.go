package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandlerPages() chi.Router {
	pagesRouter := chi.NewRouter()
	// Permite establecer la ruta html del index
	pagesRouter.Get("/index", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/index.html")
	})
	// Permite establecer la ruta html de crear usuario
	pagesRouter.Get("/createUsers", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/admin/forms/users/add_users.html")
	})
	// Permite establecer la ruta html del login
	pagesRouter.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/login/login.html")
	})
	// Permite establecer la ruta html de pagina de informacion de usuarios
	pagesRouter.Get("/usersInfo", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/users/users.html")
	})
	// Permite ingresar al formulario de actualizacion de datos usuario
	pagesRouter.Get("/FormUpdateUsers", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/admin/forms/users/updateUsers.html")
	})
	// Pagina inicial para el admin
	pagesRouter.Get("/adminHome", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/homePages/adminHome.html")
	})
	// Pagina inicial para el clientes
	pagesRouter.Get("/clientHome", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/homePages/clientHome.html")
	})
	// pagina navegador general
	pagesRouter.Get("/nav", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/components/navGeneral.html")
	})
	// Pagina informacion clientes
	pagesRouter.Get("/clientsInfo", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/admin/clients.html")
	})
	// Pagina agregar cliente
	pagesRouter.Get("/addClient", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/admin/forms/clients/add_client.html")
	})
	// Pagina para actualizar informacion cliente
	pagesRouter.Get("/FormUpdateClients", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/admin/forms/clients/updateClient.html")
	})
	// Pagina listado de regionales existentes
	pagesRouter.Get("/listRegional", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/lists/regional/regionalList.html")
	})
	// Pagina para crear regionales
	pagesRouter.Get("/formCreateRegional", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/forms/regional/regionalCreate.html")
	})
	// Pagina para actualizar informacion regionales
	pagesRouter.Get("/formUpdateRegional", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/forms/regional/regionalUpdate.html")
	})
	// pagina con menu de modulos para que usuario elija datos que desea ver segun la region
	pagesRouter.Get("/moduleMenu", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/pages/moduleMenu.html")
	})
	// pagina para moder vizualizar los contactos de cada regional
	pagesRouter.Get("/regionalContacts", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/lists/regionalContacts/regionalContacts.html")
	})
	// Pagina para creacion de nuevos contactos regional
	pagesRouter.Get("/formCreateRegionalContacts", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/forms/regionalContacts/creationForm.html")
	})
	// Pagina para actualizacion de informacion de contactos regional
	pagesRouter.Get("/formUpdateRegionalContacts", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/forms/regionalContacts/updateForm.html")
	})
	// Pagina para poder vizualizar los puntos servicios de cada regional
	pagesRouter.Get("/servicePoints", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/lists/servicePoints/servicePoints.html")
	})
	// Pagina para creacion de nuevos puntos de servicios
	pagesRouter.Get("/formCreateServicePoints", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/forms/servicePoints/creationForm.html")
	})
	// Pagina para actualizacion de nuevos puntos de servicios
	pagesRouter.Get("/formUpdateServicePoints", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/pages/clients/d1/forms/servicePoints/updateForm.html")
	})
	pagesRouter.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))

	return pagesRouter
}
