package user

import (
	"database/sql"
	"fmt"

	manager_error "github.com/LuizFreitas225/user-manager-go/src/configuration/rest_erro"
	"github.com/LuizFreitas225/user-manager-go/src/controller/user/data"
	"github.com/LuizFreitas225/user-manager-go/src/model"
)

type UserRepository struct {
	Db *sql.DB
}

func (ur *UserRepository) Search(input data.InputUserOfSearch) ([]data.OutputUserOfSearch, error) {
	users := []data.OutputUserOfSearch{}
	args := []interface{}{}

	query := fmt.Sprintf("SELECT id, create_date, last_modified_date, name, email FROM %s ", model.UserTableName)

	if input.SearchTerm != "" {
		query += " WHERE unaccent(upper(email)) LIKE unaccent(upper($1))" +
			" OR unaccent(upper(name)) LIKE unaccent(upper($2))"

		completeTerm := "%" + input.SearchTerm + "%"
		args = append(args, completeTerm, completeTerm)
	}

	rows, err := ur.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user data.OutputUserOfSearch
		err := rows.Scan(
			&user.ID,
			&user.CreateDate,
			&user.LastModifiedDate,
			&user.Name,
			&user.Email,
		)
		if err != nil {
			return nil, manager_error.NewInternalServerError("Failed to query database.", []string{err.Error()})
		}
		users = append(users, user)
	}

	// Verifica se houve erro durante a iteração
	if err := rows.Err(); err != nil {
		return nil, manager_error.NewInternalServerError("Failed to query database.", []string{err.Error()})
	}

	return users, nil
}

func (ur *UserRepository) FindById(input data.InputUserOfFindById) (data.OutputUserOfFindById, error) {

	query := fmt.Sprintf("SELECT id, create_date, last_modified_date, name, email FROM %s WHERE id = $1", model.UserTableName)
	row := ur.Db.QueryRow(query, input.ID)

	var user data.OutputUserOfFindById
	err := row.Scan(
		&user.ID,
		&user.CreateDate,
		&user.LastModifiedDate,
		&user.Name,
		&user.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return data.OutputUserOfFindById{}, manager_error.NewBadRequestError("User Not Found")
		}
		return data.OutputUserOfFindById{}, manager_error.NewInternalServerError("Failed to query database.", []string{err.Error()})
	}

	return user, nil
}

func (ur *UserRepository) Delete(input data.InputUserOfDelete) error {
	// Executa o delete diretamente
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", model.UserTableName)

	result, err := ur.Db.Exec(query, input.ID)
	if err != nil {
		return manager_error.NewInternalServerError("Failed to delete user.", []string{err.Error()})
	}

	// Verifica se algum registro foi deletado
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return manager_error.NewInternalServerError("Failed to get affected rows.", []string{err.Error()})
	}

	if rowsAffected == 0 {
		return manager_error.NewBadRequestError("User not found")
	}

	return nil
}

func (ur *UserRepository) Create(input data.InputUserOfCreate) error {

	user, err := ur.findByEmail(input.Email)

	if err != nil {
		return err
	}
    if user.ID != 0 {
     	return manager_error.NewBadRequestError("Email is already in use.")
    }

	// Executa o delete diretamente
	query := fmt.Sprintf("INSERT INTO %s (name, email, password) values($id, $1, $2, $3)", model.UserTableName)

	result, err := ur.Db.Exec(query, input.Name, input.Email, input.Password)
	if err != nil {
		return manager_error.NewInternalServerError("Failed to create user.", []string{err.Error()})
	}

	// Verifica se algum registro foi criado
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return manager_error.NewInternalServerError("Failed to get affected rows.", []string{err.Error()})
	}

	if rowsAffected == 0 {
		return manager_error.NewBadRequestError("User not created, try again later")
	}

	return nil
} 

func (ur *UserRepository) findByEmail(email string) (data.OutputUserOfFindById, error) {

	query := fmt.Sprintf("SELECT id, create_date, last_modified_date, name, email FROM %s WHERE email = $1", model.UserTableName)
	row := ur.Db.QueryRow(query, email)

	var user data.OutputUserOfFindById
	err := row.Scan(
		&user.ID,
		&user.CreateDate,
		&user.LastModifiedDate,
		&user.Name,
		&user.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return data.OutputUserOfFindById{}, nil
		}
		return data.OutputUserOfFindById{}, manager_error.NewInternalServerError("Failed to query database.", []string{err.Error()})
	}

	return user, nil
}