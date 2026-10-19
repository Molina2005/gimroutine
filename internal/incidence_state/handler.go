package incidencestate

type HandlerIncidenceState struct {
	service *ServiceIncidenceState
}

func NewHandler(h *ServiceIncidenceState) *HandlerIncidenceState {
	return &HandlerIncidenceState{service: h}
}
