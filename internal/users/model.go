package users

// Constular informacion usuarios como cracion, actualizacion, etc
type User struct {
	Id             int
	Name           string
	Email          string
	Password       string
	IdRole         int
	Role           string
	LastName       string
	Phone          string
	Id_users_state int
	Users_state    string
}

// Struct informacion usuarios consultas de todos los datos en texto
type Users struct {
	Id          int
	Name        string
	Gmail       string
	Password    string
	EntryDate   string
	Role        string
	LastName    string
	Phone       string
	Users_state string
}

// Struct search
type SearchUsers struct {
	Id          int
	Name        string
	Gmail       string
	Password    string
	EntryDate   string
	Role        string
	LastName    string
	Phone       string
	Users_state string
}
