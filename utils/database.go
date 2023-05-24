package utils

import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"
)

func GetLatestListOrder() int {
	var res models.List
	initializers.DB.Model(&models.List{}).Order(`"order" desc`).Limit(1).Find(&res)
	return res.Order
}

func GetLatestTaskOrder(listID int) int {
	var res models.Task
	initializers.DB.Model(&models.Task{}).Where("list_id = ?", listID).Order(`"order" desc`).Limit(1).Find(&res)
	return res.Order
}