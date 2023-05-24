package controllers

import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/logic"
	"ChAMP-Backend-Final-Project/models"
	"ChAMP-Backend-Final-Project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ListCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		models.Task
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a Task
	list := models.List{Title: body.Title, Order: utils.GetLatestListOrder() + 1}

	result := initializers.DB.Create(&list) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return
	c.JSON(201, gin.H{
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
	initializers.DB.Preload(clause.Associations).First(&list, id)
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
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fix range
	if body.Order < 0 {
		body.Order = 1
	} else if x := utils.GetLatestListOrder(); body.Order > x {
		body.Order = x
	}

	// Update if change order
	logic.ListReorder(list, body.Order)

	// Update basic info
	initializers.DB.Model(&list).Updates(models.List{
		Title: body.Title,
	})

	// Return
	c.JSON(200, gin.H{
		"list": list,
	})
}

func ListDelete(c *gin.Context) {
	// Find task with id
	id := c.Param("id")
	var list models.List
	initializers.DB.First(&list, id)

	// Delete all the tasks in it
	initializers.DB.Delete(&models.Task{}, "list_id = ?", id)

	// Decrease order of tasks after this task
	initializers.DB.Model(&models.List{}).Where(`"order" BETWEEN ? AND ?`, list.Order+1, utils.GetLatestListOrder()).Update(`"order"`, gorm.Expr(`"order" - 1`))

	// Delete the list
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
