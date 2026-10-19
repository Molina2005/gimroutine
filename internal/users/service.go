package users

import (
	"errors"
	"fmt"
	"modulo/internal/login"
	"modulo/internal/utils"
	"os"
	"strconv"
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
var ErrSuperAdminAlreadyExists = errors.New("superAdmin ya existe en el sistema")

// Creacion de usuario y requirimientos a seguir
func (s *ServiceUsers) ServiceCreatetUser(name, email, password, lastName string, id_rol, phone, id_users_state int) error {
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
	// Nombre y email del usuario en mayuscula para evitar usuarios duplicados en la creacion
	LowerName := strings.ToLower(name)
	LowerEmail := strings.ToLower(email)

	// Validaion de se pemiten campos vacios
	if name == "" || email == "" || HashPassword == "" || lastName == "" || id_rol == 0 || phone == 0 || id_users_state == 0 {
		return errors.New("todos los campos son obligatorios")
	}
	return s.repo.QueryInsertUser(LowerName, LowerEmail, HashPassword, lastName, phone, id_rol, id_users_state)
}

// Servicio consultar usuario
func (s *ServiceUsers) ServiceQueryUser(id_user int) (*User, error) {
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
func (s *ServiceUsers) ServiceUpdateUserInformation(id_user int, name, email, password, apellido, phone string, id_rol, id_state int) error {
	Exist, err := s.repo.QueryuserExistsXId(id_user)
	if err != nil {
		fmt.Print("err en query service ")
		return err
	}
	if !Exist {
		return ErrUserDoesNotExists
	}
	// Si no se actializa contraseña solo actualiza los demas datos
	if password == "" {
		return s.repo.QueryUpdateUserNoPassword(id_user, name, email, apellido, phone, id_rol, id_state)
	}
	HashPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("error al hashear contraseña")
	}
	return s.repo.QueryUpdateUser(id_user, name, email, HashPassword, apellido, phone, id_rol, id_state)
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
func (s *ServiceUsers) ServiceConsultAllUsers() ([]Users, error) {
	return s.repo.QueryUsersInformation()
}

// Servicio para buscar usuario por nombre o correo
func (s *ServiceUsers) ServiceUserSearch(dataUsersSearch string) ([]SearchUsers, error) {
	return s.repo.QuerySearchUsers(dataUsersSearch)
}

// Creacion de superAdmin y requirimientos a seguir
func (s *ServiceUsers) ServiceCreateSuperAdmin() error {
	Exists, err := s.repo.QueryExistsSuperAdmin()
	if err != nil {
		return err
	}
	if Exists {
		return ErrSuperAdminAlreadyExists
	}
	// Se convierte dato que viene de .env a entero
	convId, errconv := strconv.Atoi(os.Getenv("IDROL_SUPER"))
	if errconv != nil {
		return errconv
	}
	convIdState, errconv := strconv.Atoi(os.Getenv("STATE"))
	if errconv != nil {
		return errconv
	}
	// Modelos para poder guardar informacion de variables de entorno y dejar rol fijo
	SuperAdmin := User{
		Name:           os.Getenv("NAME_SUPER"),
		Email:          os.Getenv("EMAIL_SUPER"),
		Password:       os.Getenv("PASSWORD_SUPER"),
		IdRole:         convId,
		LastName:       os.Getenv("LASTNAME_SUPER"),
		Phone:          os.Getenv("PHONE_SUPER"),
		Id_users_state: convIdState,
	}

	// Contraseña incriptada
	HashPassword, err := utils.HashPassword(SuperAdmin.Password)
	if err != nil {
		return errors.New("error al hashear contraseña")
	}
	// Nombre y email del usuario en minuscula para evitar usuarios duplicados en la creacion
	LowerName := strings.ToLower(SuperAdmin.Name)
	LowerEmail := strings.ToLower(SuperAdmin.Email)
	return s.repo.QueryCreateSuperAdmin(LowerName, LowerEmail, HashPassword,
		SuperAdmin.LastName, SuperAdmin.Phone, SuperAdmin.IdRole, SuperAdmin.Id_users_state)
}
