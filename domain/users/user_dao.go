// The only point in the application where you interact with the database

package users

import (
	"fmt"
	"github.com/AlexHusleag/bookstore_users-api/datasources/mysql/users_db"
	"github.com/AlexHusleag/bookstore_users-api/utils/date"
	"github.com/AlexHusleag/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFound(fmt.Sprintf("User %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil{
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}

	user.DateCreated = date.GetNowString()

	userDB[user.Id] = user

	return nil
}
