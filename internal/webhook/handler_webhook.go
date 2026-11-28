package webhook

type HandlerWebHook struct {
	service *ServiceWebHook
}

func NewHandler(h *ServiceWebHook) *HandlerWebHook {
	return &HandlerWebHook{service: h}
}
