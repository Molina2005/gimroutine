package exercisesusers

type HandlerExercisesUsers struct {
	servicio *ServiceExercisesUsers
}

func NewHandler(s *ServiceExercisesUsers) *HandlerExercisesUsers {
	return &HandlerExercisesUsers{servicio: s}
}
