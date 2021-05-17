// entry points for the application
package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented")
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented")
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented")
}
