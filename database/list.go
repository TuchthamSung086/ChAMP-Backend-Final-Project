package database

import (
	"ChAMP-Backend-Final-Project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// INTERFACE FACTORY

// Define all services this interface provides, the controller can call these functions only
type ListService interface {
	ListCreate(list *models.ControllerList) (*models.ControllerList, error)
	ListGetAll() ([]*models.ControllerList, error)
	ListGetById(id uint) (*models.ControllerList, error)
	ListUpdate(id uint, updateBody *models.ControllerList) (*models.ControllerList, error)
	ListDelete(id uint) (*models.ControllerList, error)
	ListDeleteAll() error
}

// Our structure, stores DB
type listService struct {
	db *gorm.DB // Database
}

// Initialize our interface for controller to use
// Return a pointer to our struct as an interface
func NewListService(db *gorm.DB) ListService {
	return &listService{db: db}
}

// SERVICES AND FUNCTIONS
func (ls *listService) ListCreate(list *models.ControllerList) (*models.ControllerList, error) {
	db := ls.db
	dbList := controllerListToList(list)
	// Fix Order
	if dbList.Order == 0 {
		dbList.Order = listGetLatestOrder(ls.db) + 1
	}

	result := db.Preload(clause.Associations).Create(dbList) // pass pointer of data to Create

	// check error
	if result.Error != nil {
		return nil, result.Error
	}

	// continue happy path
	return listToControllerList(dbList), nil
}

func (ls *listService) ListGetAll() ([]*models.ControllerList, error) {
	db := ls.db
	var lists []models.List
	result := db.Preload(clause.Associations).Find(&lists)

	if result.Error != nil {
		return nil, result.Error
	}

	return listsToControllerLists(lists), nil
}

func (ls *listService) ListGetById(id uint) (*models.ControllerList, error) {
	db := ls.db
	var list models.List
	result := db.Preload(clause.Associations).First(&list, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return listToControllerList(&list), nil
}

func (ls *listService) ListUpdate(id uint, updateBody *models.ControllerList) (*models.ControllerList, error) {
	db := ls.db
	updateBody.Order = listFixOrderRange(updateBody.Order, ls.db)
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

func (ls *listService) ListDelete(id uint) (*models.ControllerList, error) {
	// Declare DB
	db := ls.db

	// Find list with id
	var list models.List
	db.First(&list, id)

	// Delete all the tasks in it
	db.Delete(&models.Task{}, "list_id = ?", id)

	// Decrease order of lists after this list
	db.Model(&models.List{}).Where(`"order" BETWEEN ? AND ?`, list.Order+1, listGetLatestOrder(ls.db)).Update(`"order"`, gorm.Expr(`"order" - 1`))

	// Save value to return deleted list
	deletedList := listToControllerList(&list)

	// Delete the list
	result := db.Delete(&models.List{}, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return deletedList, nil
}

// Hard Delete for testing
func (ls *listService) ListDeleteAll() error {
	// Delete
	result := ls.db.Unscoped().Delete(&models.List{}, "Title LIKE ?", "%")
	if result.Error != nil {
		return result.Error
	}
	return nil
}
