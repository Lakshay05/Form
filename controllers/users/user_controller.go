package users

import (
	"net/http"
	"strconv"

	"github.com/Lakshay05/go_users_api/domain/users"
	"github.com/Lakshay05/go_users_api/services"
	"github.com/Lakshay05/go_users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("iser_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id used be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("iser_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id used be a number")
		c.JSON(err.Status, err)
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch

	user.Id = userId

	result, err := services.UpdateUser(user, isPartial)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)

}

func Delete(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("iser_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id used be a number")
		c.JSON(err.Status, err)
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
}
