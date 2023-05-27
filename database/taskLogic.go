package database

import (
	"ChAMP-Backend-Final-Project/models"

	"gorm.io/gorm"
)

func (ts *taskService) getLatestTaskOrder(listID int) int {
	var res models.Task
	ts.db.Model(&models.Task{}).Where("list_id = ?", listID).Order(`"order" desc`).Limit(1).Find(&res)
	if res.Order >= 1 {
		return res.Order
	}
	return 0
}

func (ts *taskService) controllerTaskToTask(task *models.ControllerTask) *models.Task {
	return &models.Task{
		Model:       gorm.Model{ID: task.ID},
		Title:       task.Title,
		Order:       task.Order,
		ListID:      task.ListID,
		Description: task.Description,
		DueDate:     task.DueDate,
	}
}

func (ts *taskService) taskToControllerTask(task *models.Task) *models.ControllerTask {
	return &models.ControllerTask{
		ID:          task.Model.ID,
		Title:       task.Title,
		Order:       task.Order,
		ListID:      task.ListID,
		Description: task.Description,
		DueDate:     task.DueDate,
	}
}

func (ts *taskService) taskFixOrderRange(order int, listID int) int {
	// Fix range
	if order < 0 {
		return 1
	} else if x := ts.getLatestTaskOrder(listID); order > x {
		return x
	}
	return order
}

func (ts *taskService) changeList(task models.Task, listID uint) {
	// Fix Order after our task in old list
	ts.db.Model(&models.Task{}).Where(`list_id = ? AND "order" > ?`, task.ListID, task.Order).Update(`"order"`, gorm.Expr(`"order" - 1`))
	// Put to the last Order of new list
	ts.db.Model(&task).Updates(models.Task{
		Order:  ts.getLatestTaskOrder(int(listID)) + 1,
		ListID: listID,
	})
}

func (ts *taskService) taskMoveToFront(task *models.Task, to int) {
	// Assume move from B to A
	A := to
	B := task.Order
	// +1 to all from A to B-1
	ts.db.Model(&models.Task{}).Where(`list_id = ? AND "order" BETWEEN ? AND ?`, task.ListID, A, B-1).Update(`"order"`, gorm.Expr(`"order" + 1`))
	// Set own order to A
	ts.db.Model(&task).Updates(models.Task{Order: A})
}

func (ts *taskService) taskMoveToBack(task *models.Task, to int) {
	// Assume move from A to B
	A := task.Order
	B := to
	// -1 to all from A+1 to B
	ts.db.Model(&models.Task{}).Where(`list_id = ? AND "order" BETWEEN ? AND ?`, task.ListID, A+1, B).Update(`"order"`, gorm.Expr(`"order" - 1`))
	// Set own order to B
	ts.db.Model(&task).Updates(models.Task{Order: B})
}

func (ts *taskService) taskReorder(task *models.Task, to int) {
	if to != 0 && to < task.Order {
		ts.taskMoveToFront(task, to)
	} else if to != 0 && to > task.Order {
		ts.taskMoveToBack(task, to)
	}
}
