// Entire business logic of the appliction

package services

import (
	"github.com/AlexHusleag/bookstore_users-api/domain/users"
	"github.com/AlexHusleag/bookstore_users-api/utils/crypto"
	"github.com/AlexHusleag/bookstore_users-api/utils/date"
	"github.com/AlexHusleag/bookstore_users-api/utils/errors"
	"regexp"
)

const (
	emailRegex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(users.User, bool) (*users.User, *errors.RestErr)
	DeleteUser(userId int64) (*users.User, *errors.RestErr)
	SearchUsers(status string) (users.Users , *errors.RestErr)
}


func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date.GetNowString()
	user.Password = crypto.GetSHA(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) UpdateUser(user users.User, isPatch bool) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		re := regexp.MustCompile(emailRegex)
		if isPatch && (re.MatchString(current.Email) == true && user.Email == "") {
			goto skip
		} else {
			return nil, err
		}
	}

skip:
	if isPatch {
		if user.FirstName != "" {
			current.LastName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *usersService)DeleteUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Delete(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService)SearchUsers(status string) (users.Users , *errors.RestErr){
	user := &users.User{}
	return user.FindByStatus(status)
}
