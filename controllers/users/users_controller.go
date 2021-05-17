// entry points for the application
package users

import (
	"net/http"

	"github.com/adamszpilewicz/bookstore_users-api/domain/users"
	"github.com/adamszpilewicz/bookstore_users-api/services"
	"github.com/adamszpilewicz/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// create template for body
	var user users.User
	
	// take body of the request
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	// serve body of a request to a service
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented")
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented")
}
