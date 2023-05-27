package main

import (
	"ChAMP-Backend-Final-Project/controllers"
	"ChAMP-Backend-Final-Project/database"
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/routes"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	// Initialize database connection
	db, err := initializers.ConnectToDB()
	if err != nil {
		panic("Can't connect to database")
	}

	// Create the services (Interface)
	listService := database.NewListService(db)
	taskService := database.NewTaskService(db)

	// Create the controllers (Controller Object)
	listController := controllers.NewListController(listService)
	taskController := controllers.NewTaskController(taskService)

	// Setup router
	r := routes.SetupRouter()

	// Match controllers to routes
	routes.SetListControllerRoutes(r, &listController)
	routes.SetTaskControllerRoutes(r, &taskController)

	r.Run() // listen and serve on localhost:PORT
}
