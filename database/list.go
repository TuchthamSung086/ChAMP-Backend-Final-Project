package database

import (
	"ChAMP-Backend-Final-Project/models"

	"gorm.io/gorm/clause"
)

func ListCreate(list *models.ControllerList) (*models.ControllerList, error) {
	dbList := models.List{Title: list.Title, Order: list.Order}
	// Fix Order
	if dbList.Order == 0 {
		dbList.Order = listGetLatestOrder() + 1
	}

	result := db.Preload(clause.Associations).Create(&dbList) // pass pointer of data to Create

	// check error
	if result.Error != nil {
		return nil, result.Error
	}

	// continue happy path
	return listToControllerList(&dbList), nil
}

func ListGetAll() ([]*models.ControllerList, error) {
	var lists []models.List
	result := db.Preload(clause.Associations).Find(&lists)

	if result.Error != nil {
		return nil, result.Error
	}

	return listsToControllerLists(lists), nil
}

func ListGetById(id uint) *models.ControllerList {
	var list models.List
	db.Preload(clause.Associations).First(&list, id)
	return listToControllerList(&list)
}
