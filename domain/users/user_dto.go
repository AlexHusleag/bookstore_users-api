package users

import (
	"github.com/AlexHusleag/bookstore_users-api/utils/errors"
	"regexp"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"-"`
}

type Users []User

const (
	EmailRegex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	re := regexp.MustCompile(EmailRegex)

	if user.Email == "" || re.MatchString(user.Email) == false {
		return errors.NewBadRequestError("Invalid Email Address")
	}

	if user.Password == "" {
		return errors.NewBadRequestError("Invalid Password")
	}
	return nil
}
