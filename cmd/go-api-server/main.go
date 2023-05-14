package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshiyoshiharu/go-api-server/model/repository"
)

func main() {
	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUsers)

	router.Run(":8080")
}

type user struct {
	ID string `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

var users []user

func getUsers(c *gin.Context) {
	users, err := repository.NewUserRepository().GetUsers()
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, u := range users {
		if u.ID == id {
			c.IndentedJSON(http.StatusOK, u)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func postUsers(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
