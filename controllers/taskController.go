package controllers

import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"
	"ChAMP-Backend-Final-Project/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TaskCreate(c *gin.Context) {
	// Get List ID
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(400, gin.H{"msg": err})
		return
	}

	// Get data off request body
	var body struct {
		models.Task
	}
	c.Bind(&body)

	// Create a Task
	task := models.Task{Title: body.Title, Description: body.Description, DueDate: body.DueDate, Order: utils.GetLatestTaskOrder(int(id)) + 1, ListID: uint(id)}

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

	// Get data from req body
	var body struct {
		models.Task
	}
	c.Bind(&body)

	// Update
	initializers.DB.Model(&task).Updates(models.Task{
		Title:       body.Title,
		Description: body.Description,
		DueDate:     body.DueDate,
		Order:       body.Order,
	})

	// Return
	c.JSON(200, gin.H{
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
