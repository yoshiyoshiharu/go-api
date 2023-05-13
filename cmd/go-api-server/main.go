package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/yoshiyoshiharu/go-api-server/db"
	"net/http"
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
	users = []user{
		{ID: "1", Email: "test1@test.com", FirstName: "Test1", LastName: "User1"},
		{ID: "2", Email: "test2@test.com", FirstName: "Test2", LastName: "User2"},
		{ID: "3", Email: "test3@test.com", FirstName: "Test3", LastName: "User3"},
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
