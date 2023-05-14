package controller

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", GetUsers)
	router.GET("/users/:id", GetUserByID)
	router.POST("/users", PostUsers)
	router.PUT("/users/:id", UpdateUser)
	router.PATCH("/users/:id", UpdateUser)
	router.DELETE("/users/:id", DeleteUser)

	return router
}
