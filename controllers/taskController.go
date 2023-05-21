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
	//
	//initializers.DB.Save(&task)
}
