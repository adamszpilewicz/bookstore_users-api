// data access object - entry point for database (one and only)
package users

import (
	"fmt"
	"log"
	"strings"

	"github.com/adamszpilewicz/bookstore_users-api/datasources/postgres/users_db"
	"github.com/adamszpilewicz/bookstore_users-api/utils/date_utils"
	"github.com/adamszpilewicz/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_key"
	queryInsertUser = ("INSERT INTO users (first_name, last_name, email, date_created) VALUES ($1, $2, $3, $4);")
)

var (
	userDB = make(map[int64]*User)
)

// Get method on the User struct for getting in database
func (u *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := userDB[u.Id]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user with the id of %d not found", u.Id))
	}

	u.Id = result.Id
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email

	return nil
}

// Save method on the User struct for saving in database
func (u *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, date_utils.GetNowString())
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", u.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error while trying saving user to database: %s", err.Error()))
	}
	log.Println(insertResult)

	currentUser := userDB[u.Id]
	if currentUser != nil {
		if currentUser.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("mail %s already taken", u.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user with the id of %d already created", u.Id))
	}
	// u.DateCreated = date_utils.GetNowString()
	userDB[u.Id] = u
	return nil
}
