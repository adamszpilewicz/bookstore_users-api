package users

import "github.com/adamszpilewicz/bookstore_users-api/utils/errors"

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// Validate method on the User struct for validation
func (u *User) Validate() *errors.RestErr {
	if u.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
