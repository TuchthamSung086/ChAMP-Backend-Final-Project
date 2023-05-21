package controllers

import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"

	"github.com/gin-gonic/gin"
)

func TaskCreate(c *gin.Context) {
	// Get data off request body

	// Create a Task
	task := models.Task{Title: "TaskTitle", Description: "TaskDesc", DueDate: "NOWHAHA", Order: 69}

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