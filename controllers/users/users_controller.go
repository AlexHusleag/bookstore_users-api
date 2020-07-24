// app entry point, ALWAYS
// take the request, validate and send the handling to the service
// requests handled by the controller
// provide the endpoints to interacts against the users API

package users

import (
	"github.com/AlexHusleag/bookstore_users-api/domain/users"
	"github.com/AlexHusleag/bookstore_users-api/services"
	"github.com/AlexHusleag/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil { // face toata treaba de mai jos scrisa verbose
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		//TODO: return bad request to the caller
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {

	userId, userIdErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userIdErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, userErr := services.GetUser(userId)
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
