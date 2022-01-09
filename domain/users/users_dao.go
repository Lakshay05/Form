package users

import (
	users_db "github.com/Lakshay05/go_users_api/datasource/mysqlusers_db"
	"github.com/Lakshay05/go_users_api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, password) VALUES(?, ?, ?, ?)"
	queryGetUser    = "SELECT id, first_name, last_name, email from users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServeError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email); getErr != nil {
		return errors.NewInternalServeError("database error")
	}
	return nil

}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServeError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if saveErr != nil {
		return errors.NewInternalServeError("database error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServeError("database error")

	}
	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServeError("database error")
	}

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.NewInternalServeError("database error")
	}
	return nil
}
func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServeError("database error")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Id); err != nil {
		return errors.NewInternalServeError("databse error")
	}
	return nil
}
