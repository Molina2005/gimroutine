package regional

// Struct para manejo de campos de cierta regional [no para llamar todas las regionales]
type DataRegional struct {
	Name        string
	IdClient    int
	Client      string
	IdUserState int
	Users_state string
}

type AllRegionales struct {
	Id     int
	Name   string
	Client string
	State  string
}

type RegionalSearch struct {
	Id     int
	Name   string
	Client string
	State  string
}
