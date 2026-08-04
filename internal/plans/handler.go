package plans

type HandlerPlans struct {
	S *ServicePlans
}

func NewHandler(h *ServicePlans) *HandlerPlans {
	return &HandlerPlans{S: h}
}
