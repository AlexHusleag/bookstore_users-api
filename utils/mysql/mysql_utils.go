package mysql

import (
	"database/sql"
	"fmt"
	"github.com/AlexHusleag/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		switch err {
		case sql.ErrNoRows:
			return errors.NewNotFoundError("No record matching the given id")
		default:
			return errors.NewInternalServerError(
				fmt.Sprintf("Error parsing database response"))
		}
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Duplicated data")
	}

	return errors.NewInternalServerError("Error processing request")
}
