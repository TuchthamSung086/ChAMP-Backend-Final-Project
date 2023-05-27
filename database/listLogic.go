package database

import (
	"ChAMP-Backend-Final-Project/models"

	"gorm.io/gorm"
)

func (ls *listService) getLatestListOrder() int {
	var res models.List
	ls.db.Model(&models.List{}).Order(`"order" desc`).Limit(1).Find(&res)
	if res.Order >= 1 {
		return res.Order
	}
	return 0
}

func (ls *listService) taskToControllerTask(task *models.Task) *models.ControllerTask {
	return &models.ControllerTask{
		ID:          task.Model.ID,
		Title:       task.Title,
		Order:       task.Order,
		ListID:      task.ListID,
		Description: task.Description,
		DueDate:     task.DueDate,
	}
}

func (ls *listService) tasksToControllerTasks(tasks []*models.Task) []*models.ControllerTask {
	var controllerTasks = []*models.ControllerTask{}
	for _, task := range tasks {
		controllerTasks = append(controllerTasks, ls.taskToControllerTask(task))
	}
	return controllerTasks
}

func (ls *listService) controllerTaskToTask(task *models.ControllerTask) *models.Task {
	return &models.Task{
		Model:       gorm.Model{ID: task.ID},
		Title:       task.Title,
		Order:       task.Order,
		ListID:      task.ListID,
		Description: task.Description,
		DueDate:     task.DueDate,
	}
}

func (ls *listService) controllerTasksToTasks(controllerTasks []*models.ControllerTask) []*models.Task {
	var tasks = []*models.Task{}
	for _, controllerTask := range controllerTasks {
		tasks = append(tasks, ls.controllerTaskToTask(controllerTask))
	}
	return tasks
}

func (ls *listService) controllerListToList(list *models.ControllerList) *models.List {
	return &models.List{
		Model: gorm.Model{ID: list.ID},
		Title: list.Title,
		Order: list.Order,
		Tasks: ls.controllerTasksToTasks(list.Tasks),
	}
}

func (ls *listService) listToControllerList(list *models.List) *models.ControllerList {
	return &models.ControllerList{
		ID:    list.ID,
		Title: list.Title,
		Order: list.Order,
		Tasks: ls.tasksToControllerTasks(list.Tasks),
	}
}

func (ls *listService) listsToControllerLists(lists []models.List) []*models.ControllerList {
	var controllerLists = []*models.ControllerList{}
	for _, list := range lists {
		controllerLists = append(controllerLists, ls.listToControllerList(&list))
	}
	return controllerLists
}

func (ls *listService) fixListOrderRange(order int) int {
	// Fix range
	if order < 0 {
		return 1
	} else if x := ls.getLatestListOrder(); order > x {
		return x
	}
	return order
}

func (ls *listService) listMoveToFront(list *models.List, to int) {
	// Assume move from B to A
	A := to
	B := list.Order
	// +1 to all from A to B-1
	ls.db.Model(&models.List{}).Where(`"order" BETWEEN ? AND ?`, A, B-1).Update(`"order"`, gorm.Expr(`"order" + 1`))
	// Set own order to A
	ls.db.Model(&list).Updates(models.List{Order: A})
}

func (ls *listService) listMoveToBack(list *models.List, to int) {
	// Assume move from A to B
	A := list.Order
	B := to
	// -1 to all from A+1 to B
	ls.db.Model(&models.List{}).Where(`"order" BETWEEN ? AND ?`, A+1, B).Update(`"order"`, gorm.Expr(`"order" - 1`))
	// Set own order to B
	ls.db.Model(&list).Updates(models.List{Order: B})
}

func (ls *listService) listReorder(list *models.List, to int) {
	if to != 0 && to < list.Order {
		ls.listMoveToFront(list, to)
	} else if to != 0 && to > list.Order {
		ls.listMoveToBack(list, to)
	}
}
