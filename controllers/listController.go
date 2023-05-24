package controllers

import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func getLatestListOrder() int {
	var res models.List
	initializers.DB.Model(&models.List{}).Order(`"order" desc`).Limit(1).Find(&res)
	return res.Order
}

func ListCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		models.Task
	}
	c.Bind(&body)

	// Create a Task
	list := models.List{Title: body.Title, Order: getLatestListOrder() + 1}

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
	initializers.DB.Preload(clause.Associations).Find(&lists)

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

func ListDelete(c *gin.Context) {
	// Find task with id
	id := c.Param("id")

	// Delete
	initializers.DB.Delete(&models.List{}, id)

	// Return
	c.Status(200)
}

// Hard Delete for testing
func ListDeleteAll(c *gin.Context) {
	// Delete
	initializers.DB.Unscoped().Delete(&models.List{}, "Title LIKE ?", "%")

	// Return
	c.Status(200)
}
