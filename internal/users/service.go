package users

import (
	"errors"
	"fmt"
	"modulo/internal/login"
	"modulo/internal/models"
	"modulo/internal/utils"
	"strings"
)

// Recibe la conexion real db de Repository
type ServiceUsers struct {
	repo      *RepositoryUsers
	RepoLogin *login.RepoLogin
}

// Creacion nuevo servicio el cual va a guardar toda la logica interna
func NewService(r *RepositoryUsers, rl *login.RepoLogin) *ServiceUsers {
	return &ServiceUsers{repo: r, RepoLogin: rl}
}

// Variables que guarda el error de existencia de usuario
var ErrUserAlreadyExists = errors.New("usuario ya existe en el sistema")
var ErrUserDoesNotExists = errors.New("usuario no existe en el sistema")

// Creacion de usuario y requirimientos a seguir
func (s *ServiceUsers) ServiceCreatetUser(name, email, password, role string) error {
	existGmailClient, err := s.RepoLogin.QueryClientExistsByGmail(email)
	if err != nil {
		return err
	}
	existGmailUser, err := s.RepoLogin.QueryuserExistsXEmail(email)
	if err != nil {
		return err
	}
	if existGmailClient || existGmailUser {
		return ErrUserAlreadyExists
	}
	// Contraseña incriptada
	HashPassword, err := utils.HashPassword(password)
	if err != nil {
		return errors.New("error al hashear contraseña")
	}
	// Nombre y email del usuario en minuscula para evitar usuarios duplicados en la creacion
	LowerName := strings.ToLower(name)
	LowerEmail := strings.ToLower(email)

	// Validaion de se pemiten campos vacios
	if name == "" || email == "" || HashPassword == "" {
		return errors.New("todos los campos son obligatorios")
	}
	return s.repo.QueryInsertUser(LowerName, LowerEmail, HashPassword, role)
}

// Servicio consultar usuario
func (s *ServiceUsers) ServiceQueryUser(id_user int) (*models.User, error) {
	Exist, err := s.repo.QueryuserExistsXId(id_user)
	if err != nil {
		return nil, err
	}
	if !Exist {
		return nil, ErrUserDoesNotExists
	}
	return s.repo.QueryViewUserInfomation(id_user)
}

// Requerimientos actualizacion informacion usuario
func (s *ServiceUsers) ServiceUpdateUserInformation(id_usuarios int, name, email, password, role string) error {
	Exist, err := s.repo.QueryuserExistsXId(id_usuarios)
	if err != nil {
		fmt.Print("err en query service ")
		return err
	}
	if !Exist {
		return ErrUserDoesNotExists
	}
	// Si no se actializa contraseña solo actualiza los demas datos
	if password == "" {
		return s.repo.QueryUpdateUserNoPassword(id_usuarios, name, email, role)
	}
	HashPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("error al hashear contraseña")
	}
	return s.repo.QueryUpdateUser(id_usuarios, name, email, HashPassword, role)
}

// Requerimientos eliminacion usuario
func (s *ServiceUsers) ServiceDeleteUser(id_usuarios int) error {
	Exist, err := s.repo.QueryuserExistsXId(id_usuarios)
	if err != nil {
		return err
	}
	if !Exist {
		return ErrUserDoesNotExists
	}
	return s.repo.QueryDeleteUser(id_usuarios)
}

// Servicio para consultar todos los usuarios
func (s *ServiceUsers) ServiceConsultAllUsers() ([]models.Users, error) {
	return s.repo.QueryUsersInformation()
}

// Servicio para buscar usuario por nombre o correo
func (s *ServiceUsers) ServiceUserSearch(dataUsersSearch string) ([]models.SearchUsers, error) {
	return s.repo.QuerySearchUsers(dataUsersSearch)
}
