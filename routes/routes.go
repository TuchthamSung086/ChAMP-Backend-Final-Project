package routes

import (
	"ChAMP-Backend-Final-Project/controllers"

	"github.com/gin-gonic/gin"

	"ChAMP-Backend-Final-Project/docs"

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

	// r.POST("/task", controllers.TaskCreate)
	// r.GET("/tasks", controllers.TaskGetAll)
	// r.GET("/task/:id", controllers.TaskGet)
	// r.PUT("/task/:id", controllers.TaskUpdate)
	// r.DELETE("/task/:id", controllers.TaskDelete)

	// r.DELETE("/dev/clearDB", controllers.ClearDatabase)

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

// r here is router
func SetListControllerRoutes(r *gin.Engine, lc *controllers.ListController) {
	r.POST("/list", lc.Create)
	r.GET("/lists", lc.GetAll)
	r.GET("/list/:id", lc.Get)
	r.PUT("/list/:id", lc.Update)
	r.DELETE("/list/:id", lc.Delete)
}

// r here is router
func SetTaskControllerRoutes(r *gin.Engine, tc *controllers.TaskController) {
	r.POST("/task", tc.Create)
	r.GET("/tasks", tc.GetAll)
	r.GET("/task/:id", tc.GetById)
}
