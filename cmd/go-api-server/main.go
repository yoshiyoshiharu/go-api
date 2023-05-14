package main

import (
	"github.com/yoshiyoshiharu/go-api-server/controller"
)

func main() {
	router := controller.GetRouter()
	router.Run()
}
