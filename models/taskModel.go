package models

import (
	"time"

	"gorm.io/gorm"
)

/*
- Description (optional)
- Due date (optional)
- Order
*/

type Task struct {
	ID          uint `gorm:"primaryKey;autoIncrement;"`
	gorm.Model       // https://gorm.io/docs/models.html
	ListID      uint
	Title       string    `gorm:"not null;"`
	Description string    `gorm:"size:255"`
	DueDate     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Order       int       `gorm:"not null;"`
}
