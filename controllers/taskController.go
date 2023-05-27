package controllers

import (
	"ChAMP-Backend-Final-Project/database"
	"ChAMP-Backend-Final-Project/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FACTORY
type TaskController struct {
	ts database.TaskService
}

// Receive real interface (not pointer)
// Return real controller object struct (not pointer)
func NewTaskController(ts database.TaskService) TaskController {
	return TaskController{ts: ts}
}

// @Summary Create a Task
// @Schemes
// @Description Create a Task with auto-set Order (set as last Order in database) in a list specified by listID.
// @Tags Task
// @Accept json
// @Produce json
// @Param task body models.SwaggerInputCreateTask true "Details of this Task"
// @Success 200 {object} models.SwaggerTask
// @Router /task [post]
func (tc *TaskController) Create(c *gin.Context) {
	// Get data off request body
	body := &models.ControllerTask{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Task
	task, err := tc.ts.Create(body)

	if err != nil {
		//c.Status(400)
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	// Return
	c.JSON(201, gin.H{
		"task": task,
	})
}

// @Summary Get All Tasks in database
// @Schemes
// @Description Get All Tasks in database
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} models.SwaggerTasks
// @Router /tasks [get]
func (tc *TaskController) GetAll(c *gin.Context) {
	// Get all records
	tasks, err := tc.ts.GetAll()

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Return
	c.JSON(200, gin.H{
		"tasks": tasks,
	})
}

// @Summary Get Task By ID
// @Schemes
// @Description Get a task by id
// @Tags Task
// @Param id path int true "ID of task to get"
// @Accept json
// @Produce json
// @Success 200 {object} models.SwaggerTask
// @Router /task/{id} [get]
func (tc *TaskController) GetById(c *gin.Context) {
	// Find task with id
	id := c.Param("id")
	taskId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	// Get by Id from database
	task, err := tc.ts.GetById(uint(taskId))

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return
	c.JSON(200, gin.H{
		"task": task,
	})

}

// @Summary Update task by id
// @Schemes
// @Description Update task with id. Fields [Title, Order, ListID] can be updated. Changing the order will affect other tasks too, just like inserting in c++ vector. Changing list without specifying Order will put it in the last order.
// @Tags Task
// @Param id path int true "ID of task to update"
// @Param task body models.SwaggerInputUpdateTask false "Details to update"
// @Accept json
// @Produce json
// @Success 200 {object} models.SwaggerTask
// @Router /task/{id} [put]
func (tc *TaskController) Update(c *gin.Context) {
	// Find task with id
	id := c.Param("id")
	taskId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	// Get data from req body
	body := &models.ControllerTask{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.ts.Update(uint(taskId), body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return
	c.JSON(200, gin.H{
		"task": task,
	})
}

// @Summary Delete task by id
// @Schemes
// @Description Delete task with id. The orders of other tasks will be updated.
// @Tags Task
// @Param id path int true "ID of task to delete"
// @Accept json
// @Produce json
// @Success 200 {object} models.SwaggerTask
// @Router /task/{id} [delete]
func (tc *TaskController) Delete(c *gin.Context) {
	// Find task with id
	id := c.Param("id")
	taskId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	deletedTask, err := tc.ts.Delete(uint(taskId))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Return
	c.JSON(200, gin.H{"deletedTask": deletedTask})
}

/*
// Restore deleted task
func TaskRestore(c *gin.Context) {
	// Find task with id
	id := c.Param("id")

	// Restore
	initializers.DB.Unscoped().Model(&models.Task{}).Where("ID", id).Update("DeletedAt", nil)

	// Return
	c.Status(200)
}

// Hard Delete for testing
func TaskDeleteAll(c *gin.Context) {

	// Delete all records from the "tasks" table
	result := initializers.DB.Unscoped().Delete(&models.Task{}, "Title LIKE ?", "%")
	if result.Error != nil {
		fmt.Println("Failed to delete records:", result.Error)
		return
	}

	// Print the number of deleted records
	fmt.Println("Number of deleted records:", result.RowsAffected)

	// Return
	c.Status(200)
}
*/
