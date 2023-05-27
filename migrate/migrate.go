package main

import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"
)

func init() {
	initializers.LoadEnvVariables()

}

func main() {
	db, err := initializers.ConnectToDB()
	if err != nil {
		panic("Cannot migrate database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.List{})
}
