package postrgesql_utils

import (
	"strings"

	"github.com/adamszpilewicz/bookstore_users-api/utils/errors"
	"github.com/lib/pq"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*pq.Error)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record found")
		}
		return errors.NewInternalServerError("error while parsing database response")
	}

	switch sqlErr.Code {
	case "23505":
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error while processing request")
}
