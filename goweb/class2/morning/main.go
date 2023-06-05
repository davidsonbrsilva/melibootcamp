package main

import (
	"github.com/gin-gonic/gin"

	"class2/controllers"
)

func main() {
	router := gin.Default()

	router.POST("/products", controllers.CreateProduct)

	router.Run()
}
