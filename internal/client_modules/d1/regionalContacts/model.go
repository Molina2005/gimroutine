package regionalcontacts

type regionalContact struct {
	Name          string
	LastName      string
	Gmail         string
	Phone         string
	IdJogPosition int
	Job_position  string
	Regional      string
	IdUserState   int
	UserState     string
}

type RegionalesContacts struct {
	Id           int
	Name         string
	LastName     string
	Gmail        string
	Phone        string
	Job_position string
	Regional     string
	UserState    string
}

// ============================== //
// STRUCT PARA FUNCION DE VALIDACION CAMPOS
// ============================== //
// type ValidationExists struct {
// 	Name     string
// 	LastName string
// 	Gmail    string
// 	Phone    string
// }
