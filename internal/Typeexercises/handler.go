package typeexercises

type HandlerExercises struct {
	service *ServiceExercises
}

func NewHandler(s *ServiceExercises) *HandlerExercises {
	return &HandlerExercises{service: s}
}
