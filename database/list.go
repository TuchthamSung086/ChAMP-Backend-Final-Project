package database

import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"

	"gorm.io/gorm/clause"
)

func listGetLatestOrder() int {
	var res models.List
	initializers.DB.Model(&models.List{}).Order(`"order" desc`).Limit(1).Find(&res)
	if res.Order >= 1 {
		return res.Order
	}
	return 0
}

func ListCreate(list *models.ControllerList) (*models.ControllerList, error) {
	dbList := models.List{Title: list.Title, Order: list.Order}
	if dbList.Order == 0 {
		dbList.Order = listGetLatestOrder() + 1
	}

	result := db.Preload(clause.Associations).Create(&dbList) // pass pointer of data to Create

	// check error
	if result.Error != nil {
		return nil, result.Error
	}

	// continue happy path
	return &models.ControllerList{
		ID:    dbList.ID,
		Title: dbList.Title,
		Order: dbList.Order,
		Tasks: dbList.Tasks,
	}, nil
}
