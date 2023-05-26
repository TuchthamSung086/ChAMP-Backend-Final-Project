package database

import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"
)

func listGetLatestOrder() int {
	var res models.List
	initializers.DB.Model(&models.List{}).Order(`"order" desc`).Limit(1).Find(&res)
	if res.Order >= 1 {
		return res.Order
	}
	return 0
}

func controllerListToList(list *models.ControllerList) *models.List {
	return &models.List{
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

func listFixOrderRange(order int) int {
	// Fix range
	if order < 0 {
		return 1
	} else if x := listGetLatestOrder(); order > x {
		return x
	}
	return order
}
