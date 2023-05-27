package database

import (
	"ChAMP-Backend-Final-Project/initializers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockServices() (ListService, TaskService) {
	// Get environment variables
	initializers.LoadEnvVariables()
	// Initialize database connection
	db, err := initializers.ConnectToDB()
	if err != nil {
		panic("Can't connect to database")
	}
	return NewListService(db), NewTaskService(db)
}

// Get all function must return nothing for an empty database
func TestGetAllListEmptyDB(t *testing.T) {
	assert := assert.New(t)
	ls, ts := mockServices()

	// Clear Database
	ts.DeleteAll()
	ls.DeleteAll()

	lists, err := ls.GetAll()

	assert.Nil(err) // Must not error
	assert.Equal(len(lists), 0)
}
