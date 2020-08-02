// The only point in the application where you interact with the database

package users

import (
	"database/sql"
	"fmt"
	"github.com/AlexHusleag/bookstore_users-api/datasources/mysql/users_db"
	"github.com/AlexHusleag/bookstore_users-api/logger"
	"github.com/AlexHusleag/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, password, status) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, password, status FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	statement, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		logger.Error("Error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("Database error")
		//return errors.NewInternalServerError(err.Error())
	}

	defer checkIfClosed(statement)

	row := statement.QueryRow(&user.Id)
	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Password, &user.Status); err != nil {
		logger.Error("Error when trying to prepare get user by id", err)
		return errors.NewInternalServerError("Database error")
		//return mysql.ParseError(err)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	statement, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		logger.Error("Error when trying to save user statement", err)
		return errors.NewInternalServerError("Database error")
		//return errors.NewInternalServerError(err.Error())
	}
	defer checkIfClosed(statement)

	insertResult, err := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)

	if err != nil {
		logger.Error("Error when trying to save user", err)
		return errors.NewInternalServerError("Database error")
		//return mysql.ParseError(err)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("Error when trying to get last insert if after creating a new user", err)
		return errors.NewInternalServerError("Database error")
		//return mysql.ParseError(err)
	}
	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestErr {
	statement, err := users_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		logger.Error("Error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("Database error")
		//return errors.NewInternalServerError(err.Error())
	}
	defer checkIfClosed(statement)

	_, err = statement.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		//return mysql.ParseError(err)
		logger.Error("Error when trying to update user", err)
		return errors.NewInternalServerError("Database error")
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	statement, err := users_db.Client.Prepare(queryDeleteUser)

	if err != nil {
		logger.Error("Error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("Database error")
		//return errors.NewInternalServerError(err.Error())
	}
	defer checkIfClosed(statement)

	_, err = statement.Exec(user.Id)
	if err != nil {
		logger.Error("Error when trying to delete user", err)
		return errors.NewInternalServerError("Database error")
		//return mysql.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	statement, err := users_db.Client.Prepare(queryFindUserByStatus)

	if err != nil {
		logger.Error("Error when trying to find users by status statement", err)
		return nil, errors.NewInternalServerError("Database error")
		//return nil, errors.NewInternalServerError(err.Error())
	}
	defer checkIfClosed(statement)

	rows, err := statement.Query(status)
	if err != nil {
		logger.Error("Error when trying to find user by status", err)
		return nil, errors.NewInternalServerError("Database error")
		//return nil, mysql.ParseError(err)
	}
	defer checkIfClosed(rows)

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("Error when trying to scan user row into User struct", err)
			return nil, errors.NewInternalServerError("Database error")
			//return nil, mysql.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No user matching status %s", status))
	}

	return results, nil
}

func checkIfClosed(data interface{}) *errors.RestErr {
	switch data.(type) {
	case *sql.Stmt:
		if err := data.(*sql.Stmt).Close(); err != nil {
			return errors.NewInternalServerError(fmt.Sprintf("Failed to close %s", err.Error()))
		}
	case *sql.Rows:
		if err := data.(*sql.Rows).Close(); err != nil {
			return errors.NewInternalServerError(fmt.Sprintf("Failed to close %s", err.Error()))
		}
	}
	return nil
}
