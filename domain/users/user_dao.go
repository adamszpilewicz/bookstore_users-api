// data access object - entry point for database (one and only)
package users

import (
	"log"

	"github.com/adamszpilewicz/bookstore_users-api/datasources/postgres/users_db"
	"github.com/adamszpilewicz/bookstore_users-api/utils/date_utils"
	"github.com/adamszpilewicz/bookstore_users-api/utils/errors"
	"github.com/adamszpilewicz/bookstore_users-api/utils/postrgesql_utils"
)

const (
	indexUniqueEmail = "email_key"
	queryInsertUser  = ("INSERT INTO users (first_name, last_name, email, date_created) VALUES ($1, $2, $3, $4);")
	queryGetUser     = ("SELECT id, first_name, last_name, email, date_created FROM users WHERE id=$1")
	queryUpdateUser  = `
	UPDATE users 
		SET 
			first_name = $1,
			last_name = $2,
			email = $3
		WHERE
			id = $4`
)

var (
	userDB = make(map[int64]*User)
)

// Get method on the User struct for getting in database
func (u *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)

	if getErr := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); getErr != nil {
		return postrgesql_utils.ParseError(getErr)
	}
	return nil
}

// Save method on the User struct for saving in database
func (u *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, date_utils.GetNowString())
	if saveErr != nil {
		return postrgesql_utils.ParseError(saveErr)
	}

	log.Println(insertResult)

	currentUser := userDB[u.Id]
	if currentUser != nil {
		return postrgesql_utils.ParseError(saveErr)
	}
	// u.DateCreated = date_utils.GetNowString()
	userDB[u.Id] = u
	return nil
}

func (u *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Id)
	if err != nil {
		return postrgesql_utils.ParseError(err)
	}
	return nil

}
