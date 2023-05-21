package main

import (
	"ChAMP-Backend-Final-Project/controllers"
	"ChAMP-Backend-Final-Project/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/tasks", controllers.TaskCreate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
