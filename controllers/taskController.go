package controllers

import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/logic"
	"ChAMP-Backend-Final-Project/models"
	"ChAMP-Backend-Final-Project/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Create a Task
// @Schemes
// @Description Create a Task with auto-set Order (set as last Order in database) in a list specified by listID.
// @Tags Task
// @Accept json
// @Produce json
// @Param listID body string true "ID of the list this task is going to be placed in"
// @Param task body models.SwaggerInputCreateTask true "Title of this Task"
// @Success 200 {object} models.SwaggerTask
// @Router /task [post]
func TaskCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		models.Task
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a Task
	task := models.Task{Title: body.Title, Description: body.Description, DueDate: body.DueDate, Order: utils.GetLatestTaskOrder(int(body.ListID)) + 1, ListID: body.ListID}

	// Fix range
	if body.Order < 0 {
		body.Order = 1
	} else if x := utils.GetLatestTaskOrder(int(task.ListID)); body.Order > x {
		body.Order = x
	}

	// Update if change order
	logic.TaskReorder(task, body.Order)

	result := initializers.DB.Create(&task) // pass pointer of data to Create

	if result.Error != nil {
		//c.Status(400)
		c.JSON(400, gin.H{"msg": result.Error})
		return
	}
	// Return
	c.JSON(201, gin.H{
		"task":          task,
		"rows affected": result.RowsAffected,
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
func TaskGetAll(c *gin.Context) {
	// Get all records
	var tasks []models.Task
	initializers.DB.Find(&tasks)

	// Return
	c.JSON(200, gin.H{
		"tasks": tasks,
	})
}

// @Summary Get Task By ID
// @Schemes
// @Description Get a task by id
// @Tags Task
// @Param id path string true "ID of task to get"
// @Accept json
// @Produce json
// @Success 200 {object} models.SwaggerTask
// @Router /task/:id [get]
func TaskGet(c *gin.Context) {
	// Find task with id
	id := c.Param("id")
	var task models.Task
	initializers.DB.First(&task, id)
	// Return
	c.JSON(200, gin.H{
		"task": task,
	})

}

// @Summary Update task by id
// @Schemes
// @Description Update task with id. Fields [Title, Order, ListID] can be updated. Changing the order will affect other tasks too, just like inserting in c++ vector. Changing list without specifying Order will put it in the last order.
// @Tags Task
// @Param id path string true "ID of task to update"
// @Param task body models.SwaggerInputUpdateTask false "Details to update"
// @Accept json
// @Produce json
// @Success 200 {object} models.SwaggerTask
// @Router /task/:id [put]
func TaskUpdate(c *gin.Context) {
	// Find task with id
	id := c.Param("id")
	var task models.Task
	initializers.DB.First(&task, id)

	// Get data from req body
	var body struct {
		models.Task
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update if change list
	if body.ListID != 0 && task.ListID != body.ListID {
		logic.ChangeList(task, body.ListID)
	}

	// Fix range
	if body.Order < 0 {
		body.Order = 1
	} else if x := utils.GetLatestTaskOrder(int(task.ListID)); body.Order > x {
		body.Order = x
	}

	// Update if change order
	logic.TaskReorder(task, body.Order)

	// Update basic info
	initializers.DB.Model(&task).Updates(models.Task{
		Title:       body.Title,
		Description: body.Description,
		DueDate:     body.DueDate,
	})

	// Return
	c.JSON(200, gin.H{
		"task": task,
	})
}

// @Summary Delete task by id
// @Schemes
// @Description Delete task with id. The orders of other tasks will be updated.
// @Tags Task
// @Param id path string true "ID of task to delete"
// @Accept json
// @Produce json
// @Success 200 {object} models.SwaggerTask
// @Router /task/:id [delete]
func TaskDelete(c *gin.Context) {
	// Find task with id
	id := c.Param("id")
	var task models.Task
	initializers.DB.First(&task, id)

	// Decrease order of tasks after this task
	initializers.DB.Model(&models.Task{}).Where(`list_id = ? AND "order" BETWEEN ? AND ?`, task.ListID, task.Order+1, utils.GetLatestTaskOrder(int(task.ListID))).Update(`"order"`, gorm.Expr(`"order" - 1`))

	// Delete
	initializers.DB.Delete(&models.Task{}, id)

	// Return
	c.Status(200)
}

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
