package clients

type Client struct {
	Name    string
	Phone   string
	IdState int
	State   string
	Nit     string
	Gmail   string
}

// Struct informacion clientes
type Clients struct {
	Id        int
	Name      string
	Phone     string
	EntryDate string
	State     string
	Nit       string
	Gmail     string
}

// Struct search
type SearchClients struct {
	Id        int
	Name      string
	Phone     string
	EntryDate string
	State     string
	Nit       string
	Gmail     string
}
