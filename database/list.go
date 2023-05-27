package database

import (
	"ChAMP-Backend-Final-Project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ***** INTERFACE FACTORY *****

// Define all services this interface provides, the controller can call these functions only
type ListService interface {
	Create(list *models.ControllerList) (*models.ControllerList, error)
	GetAll() ([]*models.ControllerList, error)
	GetById(id uint) (*models.ControllerList, error)
	Update(id uint, updateBody *models.ControllerList) (*models.ControllerList, error)
	Delete(id uint) (*models.ControllerList, error)
	DeleteAll() error
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

// ***** SERVICES AND FUNCTIONS *****
func (ls *listService) Create(list *models.ControllerList) (*models.ControllerList, error) {
	db := ls.db
	dbList := ls.controllerListToList(list)
	// Fix Order
	if dbList.Order == 0 {
		dbList.Order = ls.getLatestListOrder() + 1
	}

	result := db.Preload(clause.Associations).Create(dbList) // pass pointer of data to Create

	// check error
	if result.Error != nil {
		return nil, result.Error
	}

	// continue happy path
	return ls.listToControllerList(dbList), nil
}

func (ls *listService) GetAll() ([]*models.ControllerList, error) {
	db := ls.db
	var lists []models.List
	result := db.Preload(clause.Associations).Find(&lists)

	if result.Error != nil {
		return nil, result.Error
	}

	return ls.listsToControllerLists(lists), nil
}

func (ls *listService) GetById(id uint) (*models.ControllerList, error) {
	db := ls.db
	var list models.List
	result := db.Preload(clause.Associations).First(&list, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return ls.listToControllerList(&list), nil
}

func (ls *listService) Update(id uint, updateBody *models.ControllerList) (*models.ControllerList, error) {
	db := ls.db
	updateBody.Order = ls.fixListOrderRange(updateBody.Order)
	var list *models.List
	db.Preload(clause.Associations).First(&list, id)

	// Update if change order
	ls.listReorder(list, updateBody.Order)

	// Update basic info
	result := db.Model(list).Updates(models.List{
		Title: updateBody.Title,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return ls.listToControllerList(list), nil
}

func (ls *listService) Delete(id uint) (*models.ControllerList, error) {
	// Declare DB
	db := ls.db

	// Find list with id
	var list models.List
	db.First(&list, id)

	// Delete all the tasks in it
	db.Delete(&models.Task{}, "list_id = ?", id)

	// Decrease order of lists after this list
	db.Model(&models.List{}).Where(`"order" BETWEEN ? AND ?`, list.Order+1, ls.getLatestListOrder()).Update(`"order"`, gorm.Expr(`"order" - 1`))

	// Save value to return deleted list
	deletedList := ls.listToControllerList(&list)

	// Delete the list
	result := db.Delete(&models.List{}, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return deletedList, nil
}

// Hard Delete for testing
func (ls *listService) DeleteAll() error {
	// Delete
	result := ls.db.Unscoped().Delete(&models.List{}, "Title LIKE ?", "%")
	if result.Error != nil {
		return result.Error
	}
	return nil
}
