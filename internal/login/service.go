package login

import (
	"modulo/internal/models"
	"modulo/internal/utils"
)

type ServiceLogin struct {
	r *RepoLogin
}

// Trae el repositorio de login para poder usar sun funciones internas
func NewService(s *RepoLogin) *ServiceLogin {
	return &ServiceLogin{r: s}
}

// Requerimientos login
func (s *ServiceLogin) ServiceLogin(email, password string) (*models.Login, error) {
	// Datos Email, password provenientes de consulta a la base de datos
	user, errHashUser := s.r.QueryLogin(email)
	// Si no hay error pasa a la comparacion de la contraseñas y a retornar la innformacion correcta
	if errHashUser == nil {
		// Comparacion de password de base de datos con password de inputPassword
		if err := utils.ComparePassword(user.Password, password); err != nil {
			return nil, err
		}
		return &models.Login{
			Id:       user.Id,
			Email:    email,
			Password: password,
			Type:     "users",
			Role:     user.Role,
		}, nil
	}
	client, errHashClient := s.r.QueryLoginClient(email)
	if errHashClient == nil {
		err := utils.ComparePassword(client.Password, password)
		if err != nil {
			return nil, err
		}
		return &models.Login{
			Id:       client.Id,
			Email:    client.Email,
			Password: client.Password,
			Type:     "clients",
		}, nil
	}
	return &models.Login{}, nil
}
