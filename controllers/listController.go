package controllers

import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"

	"github.com/gin-gonic/gin"
)

func ListCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		models.Task
	}
	c.Bind(&body)

	// Create a Task
	list := models.List{Title: body.Title, Order: body.Order}

	result := initializers.DB.Create(&list) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return
	c.JSON(200, gin.H{
		"list": list,
	})
}

func ListGetAll(c *gin.Context) {
	// Get all records
	var lists []models.List
	initializers.DB.Find(&lists)

	// Return
	c.JSON(200, gin.H{
		"lists": lists,
	})
}

func ListGet(c *gin.Context) {
	// Find task with id
	id := c.Param("id")
	var list models.List
	initializers.DB.First(&list, id)
	// Return
	c.JSON(200, gin.H{
		"list": list,
	})

}

func ListUpdate(c *gin.Context) {
	// Find task with id
	id := c.Param("id")
	var list models.List
	initializers.DB.First(&list, id)

	// Get data from req body
	var body struct {
		models.List
	}
	c.Bind(&body)

	// Update
	initializers.DB.Model(&list).Updates(models.List{
		Title: body.Title,
		Order: body.Order,
	})

	// Return
	c.JSON(200, gin.H{
		"list": list,
	})
}
