package database

import (
	"ChAMP-Backend-Final-Project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ***** INTERFACE FACTORY *****

// Define all services this interface provides, the controller can call these functions only
type TaskService interface {
	Create(list *models.ControllerTask) (*models.ControllerTask, error)
	GetAll() ([]*models.ControllerTask, error)
	// GetById(id uint) (*models.ControllerTask, error)
	// Update(id uint, updateBody *models.ControllerTask) (*models.ControllerTask, error)
	// Delete(id uint) (*models.ControllerTask, error)
	// DeleteAll() error
}

// Our structure, stores DB
type taskService struct {
	db *gorm.DB // Database
}

// Initialize our interface for controller to use
// Return a pointer to our struct as an interface
func NewTaskService(db *gorm.DB) TaskService {
	return &taskService{db: db}
}

// ***** SERVICES AND FUNCTIONS *****
func (ts *taskService) Create(task *models.ControllerTask) (*models.ControllerTask, error) {
	// Convert to real database model
	dbTask := ts.controllerTaskToTask(task)

	// Create a Task at the end of list (last order)
	dbTask.Order = ts.getLatestTaskOrder(int(dbTask.ListID))
	result := ts.db.Preload(clause.Associations).Create(dbTask) // pass pointer of data to Create
	if result.Error != nil {
		return nil, result.Error
	}

	// Fix new Order to be in possible range
	// Then, Update if change order
	ts.taskReorder(dbTask, ts.taskFixOrderRange(task.Order, int(task.ListID)))

	return ts.taskToControllerTask(dbTask), nil
}

func (ts *taskService) GetAll() ([]*models.ControllerTask, error) {
	// Get all records
	var tasks []models.Task
	result := ts.db.Preload(clause.Associations).Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return ts.tasksToControllerTasks(tasks), nil
}
