package logic

/*
import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"
	"ChAMP-Backend-Final-Project/utils"

	"gorm.io/gorm"
)

func ChangeList(task models.Task, listID uint) {
	// Fix Order after our task in old list
	initializers.DB.Model(&models.Task{}).Where(`list_id = ? AND "order" > ?`, task.ListID, task.Order).Update(`"order"`, gorm.Expr(`"order" - 1`))
	// Put to the last Order of new list
	initializers.DB.Model(&task).Updates(models.Task{
		Order:  utils.GetLatestTaskOrder(int(listID)) + 1,
		ListID: listID,
	})
}

func taskMoveToFront(task models.Task, to int) {
	// Assume move from B to A
	A := to
	B := task.Order
	// +1 to all from A to B-1
	initializers.DB.Model(&models.Task{}).Where(`list_id = ? AND "order" BETWEEN ? AND ?`, task.ListID, A, B-1).Update(`"order"`, gorm.Expr(`"order" + 1`))
	// Set own order to A
	initializers.DB.Model(&task).Updates(models.Task{Order: A})
}

func taskMoveToBack(task models.Task, to int) {
	// Assume move from A to B
	A := task.Order
	B := to
	// -1 to all from A+1 to B
	initializers.DB.Model(&models.Task{}).Where(`list_id = ? AND "order" BETWEEN ? AND ?`, task.ListID, A+1, B).Update(`"order"`, gorm.Expr(`"order" - 1`))
	// Set own order to B
	initializers.DB.Model(&task).Updates(models.Task{Order: B})
}

func TaskReorder(task models.Task, to int) {
	if to != 0 && to < task.Order {
		taskMoveToFront(task, to)
	} else if to != 0 && to > task.Order {
		taskMoveToBack(task, to)
	}
} */
