package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshiyoshiharu/go-api-server/model/entity"
	"github.com/yoshiyoshiharu/go-api-server/model/repository"
)

func main() {
	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUsers)
	router.PUT("/users/:id", updateUser)
	router.PATCH("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	router.Run(":8080")
}

func getUsers(c *gin.Context) {
	users, err := repository.NewUserRepository().GetUsers()
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := repository.NewUserRepository().GetUserByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	} else {
		c.IndentedJSON(http.StatusOK, user)
	}
}

func postUsers(c *gin.Context) {
	var newUser entity.UserEntity

	if err := c.BindJSON(&newUser); err != nil {
		fmt.Print(err)
		return
	}
	repository.NewUserRepository().CreateUser(newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}

func updateUser(c *gin.Context) {
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

func deleteUser(c *gin.Context) {
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
