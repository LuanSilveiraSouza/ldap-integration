package user

import (
	"database/sql"
)

type Repository interface {
	ListUserTypes() ([]UserType, error)
	Get(id int64) (User, error)
	List() ([]User, error)
}

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repo{db: db}
}

func (repo *repo) ListUserTypes() ([]UserType, error) {
	result := []UserType{}

	sqlString := "SELECT id, name FROM users_type;"

	sqlStatement, err := repo.db.Query(sqlString)
	if err != nil {
		return nil, err
	}
	defer sqlStatement.Close()

	for sqlStatement.Next() {
		item := UserType{}

		rows := []interface{}{&item.ID, &item.Name}

		err = sqlStatement.Scan(rows...)
		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}
	return result, nil
}

func (repo *repo) Get(id int64) (User, error) {
	result := User{}

	sqlString := "SELECT id, email, name, type_id FROM users WHERE id = $1"

	sqlStatement, err := repo.db.Query(sqlString, id)
	if err != nil {
		return result, err
	}
	defer sqlStatement.Close()

	rows := []interface{}{&result.ID, &result.Email, &result.Name, &result.Type}

	for sqlStatement.Next() {
		err = sqlStatement.Scan(rows...)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}

func (repo *repo) List() ([]User, error) {
	result := []User{}

	sqlString := `SELECT id, email, name, type_id FROM users`

	sqlStatement, err := repo.db.Query(sqlString)
	if err != nil {
		return nil, err
	}
	defer sqlStatement.Close()

	for sqlStatement.Next() {
		item := User{}

		rows := []interface{}{&item.ID, &item.Email, &item.Name, &item.Type}

		err = sqlStatement.Scan(rows...)
		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}
	return result, nil
}
