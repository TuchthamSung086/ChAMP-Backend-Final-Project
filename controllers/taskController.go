package controllers

import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"

	"github.com/gin-gonic/gin"
)

func TaskCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		models.Task
	}
	c.Bind(&body)

	// Create a Task
	task := models.Task{Title: body.Title, Description: body.Description, DueDate: body.DueDate, Order: body.Order}

	result := initializers.DB.Create(&task) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return
	c.JSON(200, gin.H{
		"task": task,
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
