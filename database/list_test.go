package database

import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test List Service, Test Task Service
var tls ListService
var tts TaskService

func init() {
	// Get environment variables
	initializers.LoadEnvVariables()
	// Initialize database connection
	db, err := initializers.ConnectToDB()
	if err != nil {
		panic("Can't connect to database")
	}
	tls = NewListService(db)
	tts = NewTaskService(db)
}

// Get all function must return nothing for an empty database
func TestGetAllListEmptyDB(t *testing.T) {
	assert := assert.New(t)
	// Clear Database
	tts.DeleteAll()
	tls.DeleteAll()

	lists, err := tls.GetAll()

	assert.Nil(err) // Must not error
	assert.Equal(len(lists), 0)
}

// Create 3 lists --> those lists must exist and have orders as 1 2 3
func TestCreateThreeLists(t *testing.T) {
	assert := assert.New(t)
	// ClearDB
	tts.DeleteAll()
	tls.DeleteAll()

	// Create 3 lists
	var cl *models.ControllerList
	var err error
	cl, err = tls.Create(&models.ControllerList{Title: "1"})
	assert.Equal(cl.Order, 1)
	assert.Nil(err)
	cl, err = tls.Create(&models.ControllerList{Title: "2"})
	assert.Equal(cl.Order, 2)
	assert.Nil(err)
	cl, err = tls.Create(&models.ControllerList{Title: "3"})
	assert.Equal(cl.Order, 3)
	assert.Nil(err)
}
