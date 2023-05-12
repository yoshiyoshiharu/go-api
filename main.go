package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yoshiyoshiharu/go-api-server/db"
)

func main() {
	db := db.Init()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": db.Stats(),
		})
	})

	router.Run(":8080")
}
