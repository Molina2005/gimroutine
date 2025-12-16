package exercises

type HandlerExercises struct {
	service *ServiceExercises
}

func NewHanlder(s *ServiceExercises) *HandlerExercises {
	return &HandlerExercises{service: s}
}
