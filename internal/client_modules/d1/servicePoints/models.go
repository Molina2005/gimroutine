package servicepoints

type ServicePoint struct {
	Name    string
	Addres  string
	City    string
	IdState int
}

type ServicePoints struct {
	Id     int
	Name   string
	Addres string
	City   string
	State  string
}

type SearchServicePoint struct {
	Id     int
	Name   string
	Addres string
	City   string
	State  string
}
