package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/neuron/src/controllers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// gin.SetMode(gin.ReleaseMode)

	route := gin.Default()

	//   database.Dbinstance()
	v1 := route.Group("/api/v1")
	{
		v1.POST("/", controllers.CreateTodo)

	}

	route.Run(":" + port)

}
