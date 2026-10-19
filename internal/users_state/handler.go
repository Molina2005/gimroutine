package usersstate

type HandlerUserState struct {
	service *ServiceUserState
}

func NewHandler(h *ServiceUserState) *HandlerUserState {
	return &HandlerUserState{service: h}
}
