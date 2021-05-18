// data access object - entry point for database (one and only)
package users

import (
	"fmt"

	"github.com/adamszpilewicz/bookstore_users-api/utils/date_utils"
	"github.com/adamszpilewicz/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

// Get method on the User struct for getting in database
func (u *User) Get() *errors.RestErr {
	result := userDB[u.Id]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user with the id of %d not found", u.Id))
	}

	u.Id = result.Id
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email

	rdea
	return nil
}

// Save method on the User struct for saving in database
func (u *User) Save() *errors.RestErr {
	currentUser := userDB[u.Id]
	if currentUser != nil {
		if currentUser.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("mail %s already taken", u.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user with the id of %d already created", u.Id))
	}
	u.DateCreated = date_utils.GetNowString()
	userDB[u.Id] = u
	return nil
}
