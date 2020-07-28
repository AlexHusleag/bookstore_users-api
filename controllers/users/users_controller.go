// app entry point, ALWAYS
// take the request, validate and send the handling to the service
// requests handled by the controller
// provide the endpoints to interacts against the users API

// domain -> service -> controller (MVC)

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

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
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

func UpdateUser(c *gin.Context) {
	userId, userIdErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userIdErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPatch := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(user, isPatch)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	userId, userIdErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userIdErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	_, userErr := services.GetUser(userId)
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	deleteUserErr, err := services.DeleteUser(userId)
	if err != nil {
		c.JSON(err.Status, deleteUserErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"Status": "User deleted"})
}

func Search(c *gin.Context){
	status := c.Query("status")
	users, err := services.FindUsers(status)
	if err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users)
}
