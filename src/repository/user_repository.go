package repository

import (
	"database/sql"
	"fmt"

	"github.com/LuizFreitas225/user-manager-go/src/controller/user/data"
	"github.com/LuizFreitas225/user-manager-go/src/model"
)

type UserRepository struct {
	Db *sql.DB
}

func (ur *UserRepository) Search(input data.InputUserOfSearch) (data.OutputUserOfSearch, error) {
	var user data.OutputUserOfSearch
	args := []interface{}{}

	query := fmt.Sprintf("SELECT id, create_date, last_modified_date, name, email, password FROM %s ", model.UserTableName)

	if input.SearchTerm != "" {
		query += " WHERE unaccent(upper(email)) LIKE unaccent(upper(?))" +
			" OR unaccent(upper(name)) LIKE unaccent(upper(?))"

		comleteTerm := "%" + input.SearchTerm + "%"
		args = append(args, comleteTerm, comleteTerm)
	}

	// Executa a query
	row := ur.Db.QueryRow(query, args...)
	err := row.Scan(
		&user.ID,
		&user.CreateDate,
		&user.LastModifiedDate,
		&user.Name,
		&user.Email,
	)

	return user, err
}
