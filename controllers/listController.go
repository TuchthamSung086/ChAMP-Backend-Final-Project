package controllers

import (
	"ChAMP-Backend-Final-Project/database"
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	body := &models.ControllerList{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a List
	list, err := database.ListCreate(body)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
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
	lists, err := database.ListGetAll()

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Return
	c.JSON(200, gin.H{
		"lists": lists,
	})
}

// @Summary Get List By ID
// @Schemes
// @Description Get a list by id
// @Tags List
// @Param id path int true "ID of list to get"
// @Accept json
// @Produce json
// @Success 200 {object} models.SwaggerList
// @Router /list/{id} [get]
func ListGet(c *gin.Context) {
	// Find task with id
	id := c.Param("id")
	listId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	list, err := database.ListGetById(uint(listId))

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return
	c.JSON(200, gin.H{
		"list": list,
	})

}

// @Summary Update list by id
// @Schemes
// @Description Update list with id. Fields [Title, Order] can be updated. Changing the order will affect other lists too, just like inserting in c++ vector.
// @Tags List
// @Param id path int true "ID of list to update"
// @Param list body models.SwaggerInputUpdateList false "Details to update"
// @Accept json
// @Produce json
// @Success 200 {object} models.SwaggerList
// @Router /list/{id} [put]
func ListUpdate(c *gin.Context) {
	// Find task with id
	id := c.Param("id")
	listId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	// Get data from req body
	body := &models.ControllerList{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list, err := database.ListUpdate(uint(listId), body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Return
	c.JSON(200, gin.H{
		"list": list,
	})
}

// @Summary Delete list by id
// @Schemes
// @Description Delete list with id. The orders of other lists will be updated.
// @Tags List
// @Param id path int true "ID of list to delete"
// @Accept json
// @Produce json
// @Success 200 {object} models.SwaggerList
// @Router /list/{id} [delete]
func ListDelete(c *gin.Context) {
	// Find task with id
	id := c.Param("id")
	listId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	deletedList, err := database.ListDelete(uint(listId))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Return
	c.JSON(200, gin.H{"deletedList": deletedList})
}

// Hard Delete for testing
func ListDeleteAll(c *gin.Context) {
	// Delete
	initializers.DB.Unscoped().Delete(&models.List{}, "Title LIKE ?", "%")

	// Return
	c.Status(200)
}
