package models

// Modelo para obtener los datos de la tabla del evento
type VitalIncidents struct {
	Id_incidencia_vital  int    `json:"id_incidencia_vital"`
	Id_punto_servicio    int    `json:"id_punto_servicio"`
	Id_usuarios          int    `json:"id_usuarios"`
	Incidencia           int    `json:"incidencia"`
	Equipo               string `json:"equipo"`
	Descripcion          string `json:"descripcion"`
	Oc                   int    `json:"oc"`
	Pendientes           string `json:"pendientes"`
	Fecha_registro       string `json:"fecha_registro"`
	Fecha_ejecucion      string `json:"fecha_ejecucion"`
	Id_status_incidencia int    `json:"id_status_incidencia"`
}

// Modelo para recibir el evento
type WebHookVitals struct {
	Type   string         `json:"type"`
	Table  string         `json:"table"`
	Schema string         `json:"schema"`
	Record VitalIncidents `json:"record"`
}
