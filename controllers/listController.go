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

// @Summary Create a List
// @Schemes
// @Description Create a List with auto-set Order (set as last Order in database)
// @Tags List
// @Accept json
// @Produce json
// @Param list body models.SwaggerInputCreateList true "Title of this List"
// @Success 200 {object} models.SwaggerList
// @Router /list [post]
func ListCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		models.List
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

// @Summary Get All Lists in database
// @Schemes
// @Description Get All Lists in database
// @Tags List
// @Accept json
// @Produce json
// @Success 200 {object} models.SwaggerLists
// @Router /lists [get]
func ListGetAll(c *gin.Context) {
	// Get all records
	var lists []models.List
	initializers.DB.Preload(clause.Associations).Find(&lists)

	// Return
	c.JSON(200, gin.H{
		"lists": lists,
	})
}

// @Summary Get List By ID
// @Schemes
// @Description Get a list by id
// @Tags List
// @Param id path string true "ID of list to get"
// @Accept json
// @Produce json
// @Success 200 {object} models.SwaggerList
// @Router /list/:id [get]
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

// @Summary Update list by id
// @Schemes
// @Description Update list with id. Fields [Title, Order] can be updated. Changing the order will affect other lists too, just like inserting in c++ vector.
// @Tags List
// @Param id path string true "ID of list to update"
// @Param list body models.SwaggerInputUpdateList false "Details to update"
// @Accept json
// @Produce json
// @Success 200 {object} models.SwaggerList
// @Router /list/:id [put]
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

// @Summary Delete list by id
// @Schemes
// @Description Delete list with id. The orders of other lists will be updated.
// @Tags List
// @Param id path string true "ID of list to delete"
// @Accept json
// @Produce json
// @Success 200 {object} models.SwaggerList
// @Router /list/:id [delete]
func ListDelete(c *gin.Context) {
	// Find task with id
	id := c.Param("id")
	var list models.List
	initializers.DB.First(&list, id)

	// Delete all the tasks in it
	initializers.DB.Delete(&models.Task{}, "list_id = ?", id)

	// Decrease order of lists after this list
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
