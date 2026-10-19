package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Guarda siempre la misma conexion a db
type RepositoryUsers struct {
	db *pgxpool.Pool
}

// Creacion nuevo repositroio para poder guardar la verdader conexion en la struct Repository
func NewRepository(db *pgxpool.Pool) *RepositoryUsers {
	return &RepositoryUsers{db: db}
}

// Verificacion existencia de usuario por id
// Uso: cuando ya estan en el sistema y se quiere validar que en realidad si esten por id
func (r RepositoryUsers) QueryuserExistsXId(id_user int) (bool, error) {
	ctx := context.Background()
	var existsId bool
	query := `SELECT EXISTS(SELECT 1 FROM usuarios WHERE id_usuarios=$1)`
	err := r.db.QueryRow(ctx, query, id_user).Scan(&existsId)
	if err != nil {
		return false, err
	}
	return existsId, nil
}

// Validacion de existencia de superAdmin
func (r *RepositoryUsers) QueryExistsSuperAdmin() (bool, error) {
	ctx := context.Background()
	var ExistRole bool
	query := `SELECT EXISTS(SELECT 1 FROM usuarios WHERE id_rol = 1)`
	if err := r.db.QueryRow(ctx, query).Scan(&ExistRole); err != nil {
		return false, err
	}
	return ExistRole, nil
}

// Insercion de usuarios
func (r *RepositoryUsers) QueryInsertUser(name, email, password, lastName string, id_rol, phone, id_users_state int) error {
	ctx := context.Background()
	query := `INSERT INTO usuarios (nombre, correo, contrasena, id_rol, apellido, telefono, id_estado_usuario) 
				VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(ctx, query, name, email, password, id_rol, lastName, phone, id_users_state)
	if err != nil {
		return err
	}
	return nil
}

// Consultar Informacion usuario
func (r *RepositoryUsers) QueryViewUserInfomation(id_user int) (*User, error) {
	ctx := context.Background()
	// Guarda la informacion que se envia con Scan
	var DataUser User
	query := `SELECT u.nombre, u.correo, u.contrasena, u.id_rol, r.cargo, 
		u.apellido, u.telefono, eu.id_estado_usuario, eu.nombre AS estado 
		FROM usuarios u 
		join rol r ON u.id_rol = r.id_rol
		join estados_usuarios eu ON eu.id_estado_usuario = u.id_estado_usuario WHERE id_usuarios=$1`
	// Envia informacion con scan a la variable DataUser
	if err := r.db.QueryRow(ctx, query, id_user).Scan(

		&DataUser.Name,
		&DataUser.Email,
		&DataUser.Password,
		&DataUser.IdRole,
		&DataUser.Role,
		&DataUser.LastName,
		&DataUser.Phone,
		&DataUser.Id_users_state,
		&DataUser.Users_state,
	); err != nil {
		return nil, err
	}
	// retorna la informacion o un error
	return &DataUser, nil
}

// Actualizacion informacion usuario con contraseña incluida
func (r *RepositoryUsers) QueryUpdateUser(id_user int, name, email, password, apellido, phone string, id_rol, id_state int) error {
	ctx := context.Background()
	query := `UPDATE usuarios SET nombre = $1, correo = $2, contrasena = $3, id_rol = $4 , apellido = $5, 
	telefono = $6, id_estado_usuario = $7 WHERE id_usuarios = $8`
	_, err := r.db.Exec(ctx, query, name, email, password, id_rol, apellido, phone, id_state, id_user)
	if err != nil {
		return err
	}
	return nil
}

// Actualizacion informacion usuario sin contraseña
func (r *RepositoryUsers) QueryUpdateUserNoPassword(id_user int, name, email, apellido, phone string, id_rol, id_state int) error {
	ctx := context.Background()
	query := `UPDATE usuarios SET nombre = $1, correo = $2, id_rol = $3 , apellido = $4, 
	telefono = $5, id_estado_usuario = $6 WHERE id_usuarios = $7`
	_, err := r.db.Exec(ctx, query, name, email, id_rol, apellido, phone, id_state, id_user)
	if err != nil {
		return err
	}
	return nil
}

// Eliminar usuario del sistema
func (r *RepositoryUsers) QueryDeleteUser(id_user int) error {
	ctx := context.Background()
	query := `DELETE FROM usuarios WHERE id_usuarios = $1`
	_, err := r.db.Exec(ctx, query, id_user)
	if err != nil {
		return err
	}
	return nil
}

// Consulta informacion de todos los usuarios
func (r *RepositoryUsers) QueryUsersInformation() ([]Users, error) {
	ctx := context.Background()
	query := `SELECT u.id_usuarios, u.nombre, u.correo, u.contrasena, u.fecha_ingreso, r.cargo, u.apellido, u.telefono, eu.nombre AS estado 
		FROM usuarios u 
		join rol r ON u.id_rol = r.id_rol
		join estados_usuarios eu ON eu.id_estado_usuario = u.id_estado_usuario 
		ORDER BY u.id_usuarios ASC`
	data, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var DataUsers []Users
	for data.Next() {
		var users Users
		if err := data.Scan(&users.Id, &users.Name, &users.Gmail, &users.Password, &users.EntryDate,
			&users.Role, &users.LastName, &users.Phone, &users.Users_state); err != nil {
			return nil, err
		}
		DataUsers = append(DataUsers, users)
	}
	return DataUsers, nil
}

// Buscar usuario en buscador
func (r *RepositoryUsers) QuerySearchUsers(search string) ([]SearchUsers, error) {
	ctx := context.Background()
	query := `SELECT u.id_usuarios, u.nombre, u.correo, u.contrasena, u.fecha_ingreso, r.cargo, u.apellido, u.telefono, eu.nombre AS estado 
		FROM usuarios u 
		join rol r ON u.id_rol = r.id_rol
		join estados_usuarios eu ON eu.id_estado_usuario = u.id_estado_usuario
		WHERE u.nombre ILIKE '%' || $1 || '%' 
		OR u.correo ILIKE '%' || $1 || '%'`
	data, err := r.db.Query(ctx, query, search)
	if err != nil {
		return nil, err
	}
	var DataUsers []SearchUsers
	defer data.Close()
	for data.Next() {
		var User SearchUsers
		err := data.Scan(&User.Id, &User.Name, &User.Gmail, &User.Password, &User.EntryDate, &User.Role, &User.LastName, &User.Phone, &User.Users_state)
		if err != nil {
			return nil, err
		}
		DataUsers = append(DataUsers, User)
	}
	return DataUsers, nil
}

// Creacion super admin
func (r *RepositoryUsers) QueryCreateSuperAdmin(name, email, password, lastName, phone string, id_rol, id_users_state int) error {
	ctx := context.Background()
	query := `INSERT INTO usuarios (nombre, correo, contrasena, id_rol, apellido, telefono, id_estado_usuario) 
				VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(ctx, query, name, email, password, id_rol, lastName, phone, id_users_state)
	if err != nil {
		return err
	}
	return nil
}
