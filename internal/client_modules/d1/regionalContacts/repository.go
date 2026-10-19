package regionalcontacts

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryRegionalContacts struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *RepositoryRegionalContacts {
	return &RepositoryRegionalContacts{db: db}
}

// Funcion para validar existencia por id
func (r *RepositoryRegionalContacts) QueryRegionalContactsExitsById(id int) (bool, error) {
	ctx := context.Background()
	var ExistsId bool
	query := `SELECT EXISTS (SELECT 1 FROM contactos_regional WHERE id_contactos_regional = $1 )`
	err := r.db.QueryRow(ctx, query, id).Scan(&ExistsId)
	if err != nil {
		return false, err
	}
	return ExistsId, nil
}

// Funcion para validar existencia por nombre
func (r *RepositoryRegionalContacts) QueryRegionalContactsExitsByName(name string) (bool, error) {
	ctx := context.Background()
	var ExistsName bool
	query := `SELECT EXISTS (SELECT 1 FROM contactos_regional WHERE nombre = $1 )`
	err := r.db.QueryRow(ctx, query, name).Scan(&ExistsName)
	if err != nil {
		return false, err
	}
	return ExistsName, nil
}

// Funcion para validar existencia por correo
func (r *RepositoryRegionalContacts) QueryRegionalContactsExitsByGmail(gmail string) (bool, error) {
	ctx := context.Background()
	var ExistsGmail bool
	query := `SELECT EXISTS (SELECT 1 FROM contactos_regional WHERE correo = $1)`
	err := r.db.QueryRow(ctx, query, gmail).Scan(&ExistsGmail)
	if err != nil {
		return false, err
	}
	return ExistsGmail, nil
}

// ======================================================================================== //
// HACER FUNCION PARA VALIDAR MAS ADELANTE TODO Y ASI NO TENER PROBLEMAS CON DATOS DUPLICADOS
// ======================================================================================== //
// Funcion para validacion de existencia y que datos no se repitan al momento de hacer cambios
// func (r *RepositoryRegionalContacts) QueryDataExists(id int, data ValidationExists) (bool, error) {
// 	ctx := context.Background()
// 	var Exists bool
// 	query := `SELECT EXISTS (SELECT 1 FROM contactos_regional
// 		WHERE (name = $1
// 		OR lastname = $2
// 		OR gmail = $3
// 		OR phone = $4)
// 		AND id_contactos_regional <> $5
// 		)`
// 	err := r.db.QueryRow(ctx, query, data.Name, data.LastName, data.Gmail, data.Phone, id).Scan(&Exists)
// 	if err != nil {
// 		return false, err
// 	}
// 	return Exists, nil
// }

// Funcion creacion de contactos regional
func (r *RepositoryRegionalContacts) QueryRegionalContactsCreate(
	name, lastName, gmail, phone string,
	id_job_position, id_regional, id_user_state int) error {
	ctx := context.Background()
	query := `INSERT INTO contactos_regional (nombre, apellido, correo, telefono, id_cargo, id_regional, id_estado_usuario) 
					VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(ctx, query, name, lastName, gmail, phone, id_job_position, id_regional, id_user_state)
	if err != nil {
		return err
	}
	return nil
}

// Funcion para consultar contacto de regional especifico
func (r *RepositoryRegionalContacts) QueryRegionalContactsInformation(idRegionalContact int) (*regionalContact, error) {
	ctx := context.Background()
	// Guarda la informacion que se envia con Scan
	var DataRegionalContact regionalContact
	query := `SELECT cr.nombre, cr.apellido, cr.correo, cr.telefono, r.nombre AS regional, eu.id_estado_usuario, eu.nombre AS estado, c.id_cargo, c.nombre AS cargo  
			FROM contactos_regional cr 
			JOIN regionales r ON r.id_regional = cr.id_regional
			JOIN estados_usuarios eu ON eu.id_estado_usuario = cr.id_estado_usuario
			JOIN cargos c ON c.id_cargo = cr.id_cargo 
			WHERE cr.id_contactos_regional = $1`
	// Envia informacion con scan a la variable DataRegionalContact
	if err := r.db.QueryRow(ctx, query, idRegionalContact).Scan(
		&DataRegionalContact.Name,
		&DataRegionalContact.LastName,
		&DataRegionalContact.Gmail,
		&DataRegionalContact.Phone,
		&DataRegionalContact.Regional,
		&DataRegionalContact.IdUserState,
		&DataRegionalContact.UserState,
		&DataRegionalContact.IdJogPosition,
		&DataRegionalContact.Job_position,
	); err != nil {
		return nil, err
	}
	// retorna la informacion o un error
	return &DataRegionalContact, nil
}

// Funcion para consultar todos los contactos regional
func (r *RepositoryRegionalContacts) QueryAllRegionalContacts(id_regional int) ([]RegionalesContacts, error) {
	ctx := context.Background()
	query := `SELECT cr.id_contactos_regional, cr.nombre, cr.apellido, cr.correo, cr.telefono, r.nombre AS regional, eu.nombre AS estado, c.nombre AS cargo  
			FROM contactos_regional cr 
			JOIN regionales r ON r.id_regional = cr.id_regional
			JOIN estados_usuarios eu ON eu.id_estado_usuario = cr.id_estado_usuario
			JOIN cargos c ON c.id_cargo = cr.id_cargo 
			WHERE cr.id_regional = $1`
	data, err := r.db.Query(ctx, query, id_regional)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var DataRegionalContacts []RegionalesContacts
	for data.Next() {
		var ModelsRegionalContacts RegionalesContacts
		if err := data.Scan(
			&ModelsRegionalContacts.Id, &ModelsRegionalContacts.Name,
			&ModelsRegionalContacts.LastName, &ModelsRegionalContacts.Gmail, &ModelsRegionalContacts.Phone,
			&ModelsRegionalContacts.Regional, &ModelsRegionalContacts.UserState, &ModelsRegionalContacts.Job_position); err != nil {
			return nil, err
		}
		DataRegionalContacts = append(DataRegionalContacts, ModelsRegionalContacts)
	}
	return DataRegionalContacts, nil
}

// Funcion para actualizar informacion de contactos
func (r *RepositoryRegionalContacts) QueryRegionalContactsUpdate(
	name, lastName, gmail, phone string,
	id_job_position, id_regional, id_user_state, idregionalContact int) error {
	ctx := context.Background()
	query := `UPDATE contactos_regional 
		SET nombre = $1, apellido = $2, correo = $3, telefono = $4, id_regional = $5, 
		id_estado_usuario = $6, id_cargo = $7 WHERE id_contactos_regional = $8`
	_, err := r.db.Exec(ctx, query, name, lastName, gmail, phone,
		id_regional, id_user_state, id_job_position, idregionalContact)
	if err != nil {
		return err
	}
	return nil
}

// Funcion eliminar contacto regional
func (r *RepositoryRegionalContacts) QueryRegionalContactsDelete(idRegionalContact int) error {
	ctx := context.Background()
	query := `DELETE FROM contactos_regional WHERE id_contactos_regional = $1`
	_, err := r.db.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

// Funcion buscador de contactos regional
func (r *RepositoryRegionalContacts) QueryRegionalSeach() {

}
