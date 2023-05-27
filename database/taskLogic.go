package database

import (
	"ChAMP-Backend-Final-Project/models"

	"gorm.io/gorm"
)

func controllerTaskToTask(task *models.ControllerTask) *models.Task {
	return &models.Task{
		Model:       gorm.Model{ID: task.ID},
		Title:       task.Title,
		Order:       task.Order,
		ListID:      task.ListID,
		Description: task.Description,
		DueDate:     task.DueDate,
	}
}

func taskFixOrderRange(order int, db *gorm.DB) int {
	// Fix range
	if order < 0 {
		return 1
	} else if x := listGetLatestOrder(db); order > x {
		return x
	}
	return order
}
