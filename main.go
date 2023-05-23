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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/task", controllers.TaskCreate)
	r.GET("/tasks", controllers.TaskGetAll)
	r.GET("/task/:id", controllers.TaskGet)
	r.PUT("/task/:id", controllers.TaskUpdate)
	r.DELETE("/task/:id", controllers.TaskDelete)

	r.POST("/list", controllers.ListCreate)
	r.GET("/lists", controllers.ListGetAll)
	r.GET("/list/:id", controllers.ListGet)
	r.PUT("/list/:id", controllers.ListUpdate)
	r.DELETE("/list/:id", controllers.ListDelete)
	r.Run() // listen and serve on localhost:PORT
}
