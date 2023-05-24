package controllers

import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"
	"ChAMP-Backend-Final-Project/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TaskCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		models.Task
	}
	c.Bind(&body)

	// Create a Task
	task := models.Task{Title: body.Title, Description: body.Description, DueDate: body.DueDate, Order: utils.GetLatestTaskOrder(int(body.ListID)) + 1, ListID: body.ListID}

	result := initializers.DB.Create(&task) // pass pointer of data to Create

	if result.Error != nil {
		//c.Status(400)
		c.JSON(400, gin.H{"msg": result.Error})
		return
	}
	// Return
	c.JSON(200, gin.H{
		"task":          task,
		"rows affected": result.RowsAffected,
	})
}

func TaskGetAll(c *gin.Context) {
	// Get all records
	var tasks []models.Task
	initializers.DB.Find(&tasks)

	// Return
	c.JSON(200, gin.H{
		"tasks": tasks,
	})
}

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

func TaskUpdate(c *gin.Context) {
	// Find task with id
	id := c.Param("id")
	var task models.Task
	initializers.DB.First(&task, id)

	done := ""

	// Get data from req body
	var body struct {
		models.Task
	}
	c.Bind(&body)

	// Update if change list

	if body.ListID != 0 && task.ListID != body.ListID {
		done += "list "
		// Fix Order after our task in old list
		initializers.DB.Model(&models.Task{}).Where(`list_id = ? AND "order" > ?`, task.ListID, task.Order).Update(`"order"`, gorm.Expr(`"order" - 1`))
		// Put to the last Order of new list
		initializers.DB.Model(&task).Updates(models.Task{
			Order:  utils.GetLatestTaskOrder(int(body.ListID)) + 1,
			ListID: body.ListID,
		})
	}

	// Update if change order
	// Move to before
	if body.Order < 0 {
		body.Order = 0
	} else if x := utils.GetLatestTaskOrder(int(task.ListID)); body.Order > x {
		body.Order = x + 1
	}
	if body.Order != 0 && body.Order < task.Order {
		done += "movebefore "
		// Assume move from B to A
		A := body.Order
		B := task.Order
		// +1 to all from A to B-1
		initializers.DB.Model(&models.Task{}).Where(`list_id = ? AND "order" BETWEEN ? AND ?`, task.ListID, A, B-1).Update(`"order"`, gorm.Expr(`"order" + 1`))
		// Set own order to A
		initializers.DB.Model(&task).Updates(models.Task{Order: A})
	} else if body.Order != 0 && body.Order > task.Order {
		done += "moveafter "
		// Assume move from A to B
		A := task.Order
		B := body.Order
		// -1 to all from A+1 to B
		initializers.DB.Model(&models.Task{}).Where(`list_id = ? AND "order" BETWEEN ? AND ?`, task.ListID, A+1, B).Update(`"order"`, gorm.Expr(`"order" - 1`))
		// Set own order to B
		initializers.DB.Model(&task).Updates(models.Task{Order: B})
	}

	// Update basic info
	initializers.DB.Model(&task).Updates(models.Task{
		Title:       body.Title,
		Description: body.Description,
		DueDate:     body.DueDate,
	})

	// Return
	c.JSON(200, gin.H{
		"done": done,
		"task": task,
	})
}

func TaskDelete(c *gin.Context) {
	// Find task with id
	id := c.Param("id")

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
