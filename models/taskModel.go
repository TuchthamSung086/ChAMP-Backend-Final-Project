package models

import "gorm.io/gorm"

/*
- Description (optional)
- Due date (optional)
- Order
*/

type Task struct {
	gorm.Model         // https://gorm.io/docs/models.html
	Title       string `gorm:"not null;"`
	Description string
	DueDate     string
	Order       int `gorm:"not null;"`
}
