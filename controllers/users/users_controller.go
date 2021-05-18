// entry points for the application
package users

import (
	"net/http"
	"strconv"

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
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	// serve body of a request to a service
	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented")
}
