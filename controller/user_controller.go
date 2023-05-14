package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshiyoshiharu/go-api-server/model/entity"
	"github.com/yoshiyoshiharu/go-api-server/model/repository"
)

func GetUsers(c *gin.Context) {
	users, err := repository.NewUserRepository().GetUsers()
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := repository.NewUserRepository().GetUserByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	} else {
		c.IndentedJSON(http.StatusOK, user)
	}
}

func PostUsers(c *gin.Context) {
	var newUser entity.UserEntity

	if err := c.BindJSON(&newUser); err != nil {
		fmt.Print(err)
		return
	}
	repository.NewUserRepository().CreateUser(newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user, err := repository.NewUserRepository().GetUserByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}

	if err := c.BindJSON(&user); err != nil {
		fmt.Print(err)
		return
	}

	repository.NewUserRepository().UpdateUser(user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	_, err := repository.NewUserRepository().GetUserByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}

	err = repository.NewUserRepository().DeleteUser(id)
	if err != nil {
		fmt.Print(err)
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "user deleted"})
}
