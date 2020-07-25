// The only point in the application where you interact with the database

package users

import (
	"database/sql"
	"fmt"
	"github.com/AlexHusleag/bookstore_users-api/datasources/mysql/users_db"
	"github.com/AlexHusleag/bookstore_users-api/utils/date"
	"github.com/AlexHusleag/bookstore_users-api/utils/errors"
	"strings"
)

const (
	uniqueEmail     = "email_UNIQUE"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT * FROM users WHERE id=?;"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	statement, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer checkIfDatabaseIsClosed(statement)

	row := statement.QueryRow(&user.Id)
	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if err == sql.ErrNoRows {
			return errors.NewNotFound(fmt.Sprintf("User %d not found", user.Id))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("Error when trying to get user %d: %s", user.Id, err.Error()))
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	statement, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer checkIfDatabaseIsClosed(statement)

	user.DateCreated = date.GetNowString()

	insertResult, err := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), uniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("Failed to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("Failed to save user: %s", err.Error()))
	}
	user.Id = userId

	return nil
}

func checkIfDatabaseIsClosed(statement *sql.Stmt) *errors.RestErr {
	if err := statement.Close(); err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Failed to close the user database %s", err.Error()))
	}
	return nil
}
