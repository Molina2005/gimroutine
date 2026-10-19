package jobposition

type HandlerJobPosition struct {
	service *ServiceJobPosition
}

func NewHandler(h *ServiceJobPosition) *HandlerJobPosition {
	return &HandlerJobPosition{service: h}
}
