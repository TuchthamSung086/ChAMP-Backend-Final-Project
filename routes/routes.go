package routes

import (
	"ChAMP-Backend-Final-Project/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter initializes the Gin router and registers all routes
func SetupRouter() *gin.Engine {
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

	r.DELETE("/dev/clearDB", controllers.ClearDatabase)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
