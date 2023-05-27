/* package logic

import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"

	"gorm.io/gorm"
)

func listMoveToFront(list models.List, to int) {
	// Assume move from B to A
	A := to
	B := list.Order
	// +1 to all from A to B-1
	initializers.DB.Model(&models.List{}).Where(`"order" BETWEEN ? AND ?`, A, B-1).Update(`"order"`, gorm.Expr(`"order" + 1`))
	// Set own order to A
	initializers.DB.Model(&list).Updates(models.List{Order: A})
}

func listMoveToBack(list models.List, to int) {
	// Assume move from A to B
	A := list.Order
	B := to
	// -1 to all from A+1 to B
	initializers.DB.Model(&models.List{}).Where(`"order" BETWEEN ? AND ?`, A+1, B).Update(`"order"`, gorm.Expr(`"order" - 1`))
	// Set own order to B
	initializers.DB.Model(&list).Updates(models.List{Order: B})
}

func ListReorder(list models.List, to int) {
	if to != 0 && to < list.Order {
		listMoveToFront(list, to)
	} else if to != 0 && to > list.Order {
		listMoveToBack(list, to)
	}
}

func hi() string {
	return "hi"
}
*/