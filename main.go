package main

import (
	"ChAMP-Backend-Final-Project/database"
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	database.ConnectToDB()
}

func main() {
	r := routes.SetupRouter()
	r.Run() // listen and serve on localhost:PORT
}
