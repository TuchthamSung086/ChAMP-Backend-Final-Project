package database

import (
	"ChAMP-Backend-Final-Project/models"

	"gorm.io/gorm"
)

func listGetLatestOrder(db *gorm.DB) int {
	var res models.List
	db.Model(&models.List{}).Order(`"order" desc`).Limit(1).Find(&res)
	if res.Order >= 1 {
		return res.Order
	}
	return 0
}

func controllerListToList(list *models.ControllerList) *models.List {
	return &models.List{
		Model: gorm.Model{ID: list.ID},
		Title: list.Title,
		Order: list.Order,
		Tasks: list.Tasks,
	}
}

func listToControllerList(list *models.List) *models.ControllerList {
	return &models.ControllerList{
		ID:    list.ID,
		Title: list.Title,
		Order: list.Order,
		Tasks: list.Tasks,
	}
}

func listsToControllerLists(lists []models.List) []*models.ControllerList {
	var controllerLists = []*models.ControllerList{}
	for _, list := range lists {
		controllerLists = append(controllerLists, listToControllerList(&list))
	}
	return controllerLists
}

func listFixOrderRange(order int, db *gorm.DB) int {
	// Fix range
	if order < 0 {
		return 1
	} else if x := listGetLatestOrder(db); order > x {
		return x
	}
	return order
}

func listMoveToFront(list *models.List, to int) {
	// Assume move from B to A
	A := to
	B := list.Order
	// +1 to all from A to B-1
	db.Model(&models.List{}).Where(`"order" BETWEEN ? AND ?`, A, B-1).Update(`"order"`, gorm.Expr(`"order" + 1`))
	// Set own order to A
	db.Model(&list).Updates(models.List{Order: A})
}

func listMoveToBack(list *models.List, to int) {
	// Assume move from A to B
	A := list.Order
	B := to
	// -1 to all from A+1 to B
	db.Model(&models.List{}).Where(`"order" BETWEEN ? AND ?`, A+1, B).Update(`"order"`, gorm.Expr(`"order" - 1`))
	// Set own order to B
	db.Model(&list).Updates(models.List{Order: B})
}

func listReorder(list *models.List, to int) {
	if to != 0 && to < list.Order {
		listMoveToFront(list, to)
	} else if to != 0 && to > list.Order {
		listMoveToBack(list, to)
	}
}
