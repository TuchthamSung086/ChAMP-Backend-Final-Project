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
	GetById(id uint) (*models.ControllerTask, error)
	Update(id uint, updateBody *models.ControllerTask) (*models.ControllerTask, error)
	Delete(id uint) (*models.ControllerTask, error)
	DeleteAll() (int, error)
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
	dbTask.Order = ts.getLatestTaskOrder(int(dbTask.ListID)) + 1
	result := ts.db.Preload(clause.Associations).Create(dbTask) // pass pointer of data to Create
	if result.Error != nil {
		return nil, result.Error
	}

	// Fix new Order to be in possible range
	// Then, Update if change order
	ts.taskReorder(dbTask, ts.fixTaskOrderRange(task.Order, int(task.ListID)))

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

func (ts *taskService) GetById(id uint) (*models.ControllerTask, error) {
	var task models.Task
	result := ts.db.Preload(clause.Associations).First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return ts.taskToControllerTask(&task), nil
}

func (ts *taskService) Update(id uint, updateBody *models.ControllerTask) (*models.ControllerTask, error) {
	// Get Task to Update
	var task *models.Task
	ts.db.Preload(clause.Associations).First(&task, id)

	// Update if change list
	if updateBody.ListID != 0 && task.ListID != updateBody.ListID {
		ts.changeList(task, updateBody.ListID)
	}

	// Fix order to be in possible range
	updateBody.Order = ts.fixTaskOrderRange(updateBody.Order, int(task.ListID))

	// Update if change order
	ts.taskReorder(task, updateBody.Order)

	// Update basic info
	result := ts.db.Model(task).Updates(models.Task{
		Title:       updateBody.Title,
		Description: updateBody.Description,
		DueDate:     updateBody.DueDate,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return ts.taskToControllerTask(task), nil
}

func (ts *taskService) Delete(id uint) (*models.ControllerTask, error) {
	// Find task with id
	var task models.Task
	ts.db.First(&task, id)

	// Decrease order of tasks after this task
	ts.db.Model(&models.Task{}).Where(`list_id = ? AND "order" BETWEEN ? AND ?`, task.ListID, task.Order+1, ts.getLatestTaskOrder(int(task.ListID))).Update(`"order"`, gorm.Expr(`"order" - 1`))

	// Save value to return deleted task
	deletedTask := ts.taskToControllerTask(&task)

	// Delete
	result := ts.db.Delete(&models.Task{}, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return deletedTask, nil
}

// Hard Delete for testing
func (ts *taskService) DeleteAll() (int, error) {
	// Delete
	result := ts.db.Unscoped().Delete(&models.Task{}, "Title LIKE ?", "%")
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}
