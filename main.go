package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/users", getUsers)

	router.Run(":8080")
}

type user struct {
	ID int `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func getUsers(c *gin.Context) {
	usres := []user{
		{ID: 1, Email: "test1@test.com", FirstName: "Test1", LastName: "User1"},
		{ID: 2, Email: "test2@test.com", FirstName: "Test2", LastName: "User2"},
		{ID: 3, Email: "test3@test.com", FirstName: "Test3", LastName: "User3"},
	}
	c.IndentedJSON(http.StatusOK, usres)
}
