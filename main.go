package main

import (
	"Demo/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	v := router.Group("/")
	{
		v.GET("/", controller.IndexController)
		v.GET("/go-test/trace", controller.CallJavaDemo)
	}

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
