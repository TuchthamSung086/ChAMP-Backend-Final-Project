package database

import (
	"ChAMP-Backend-Final-Project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ListCreate(list *models.ControllerList) (*models.ControllerList, error) {
	dbList := controllerListToList(list)
	// Fix Order
	if dbList.Order == 0 {
		dbList.Order = listGetLatestOrder() + 1
	}

	result := db.Preload(clause.Associations).Create(dbList) // pass pointer of data to Create

	// check error
	if result.Error != nil {
		return nil, result.Error
	}

	// continue happy path
	return listToControllerList(dbList), nil
}

func ListGetAll() ([]*models.ControllerList, error) {
	var lists []models.List
	result := db.Preload(clause.Associations).Find(&lists)

	if result.Error != nil {
		return nil, result.Error
	}

	return listsToControllerLists(lists), nil
}

func ListGetById(id uint) (*models.ControllerList, error) {
	var list models.List
	result := db.Preload(clause.Associations).First(&list, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return listToControllerList(&list), nil
}

func ListUpdate(id uint, updateBody *models.ControllerList) (*models.ControllerList, error) {
	updateBody.Order = listFixOrderRange(updateBody.Order)
	var list *models.List
	db.Preload(clause.Associations).First(&list, id)

	// Update if change order
	listReorder(list, updateBody.Order)

	// Update basic info
	result := db.Model(list).Updates(models.List{
		Title: updateBody.Title,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return listToControllerList(list), nil
}

func ListDelete(id uint) (*models.ControllerList, error) {
	// Find list with id
	var list models.List
	db.First(&list, id)

	// Delete all the tasks in it
	db.Delete(&models.Task{}, "list_id = ?", id)

	// Decrease order of lists after this list
	db.Model(&models.List{}).Where(`"order" BETWEEN ? AND ?`, list.Order+1, listGetLatestOrder()).Update(`"order"`, gorm.Expr(`"order" - 1`))

	// Save value to return deleted list
	deletedList := listToControllerList(&list)

	// Delete the list
	result := db.Delete(&models.List{}, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return deletedList, nil
}
