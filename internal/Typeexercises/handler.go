package typeexercises

type HandlerTypeOfExercises struct {
	service *ServiceTypeOfExercises
}

func NewHandler(s *ServiceTypeOfExercises) *HandlerTypeOfExercises {
	return &HandlerTypeOfExercises{service: s}
}
