package login

type HandlerLogin struct {
	S *ServiceLogin
}

func NewHandler(h *ServiceLogin) *HandlerLogin {
	return &HandlerLogin{S: h}
}
