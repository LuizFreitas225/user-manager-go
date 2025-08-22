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

	// Executa a query
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
