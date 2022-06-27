package login

import (
	"database/sql"
	user "ldap-integration/src/user"
)

type Repository interface {
	Authenticate(email string) (user.User, error)
}

type repo struct {
	db *sql.DB
}

func NewRepo(DB *sql.DB) Repository {
	return &repo{db: DB}
}

func (repo *repo) Authenticate(email string) (user.User, error) {
	result := user.User{Email: email}

	sqlString := "SELECT id, name, type_id FROM users WHERE email = $1"

	err := repo.db.QueryRow(sqlString, email).Scan(&result.ID, &result.Name, &result.Type)
	if err != nil {
		return result, err
	}

	return result, nil
}
