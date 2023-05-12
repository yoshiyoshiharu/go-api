package main

import(
	"github.com/gin-gonic/gin"
	"github.com/yoshiyoshiharu/go-api-server/db"
	"fmt"
)

func main() {
	db := db.Init()
	fmt.Printf("%+v", db.Stats())

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.Run(":8080")
}
