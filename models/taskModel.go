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

// For database
type Task struct {
	gorm.Model            // https://gorm.io/docs/models.html
	ListID      uint      `gorm:"not null;"`
	Title       string    `gorm:"not null;size:255;default:Untitled;"`
	Description string    `gorm:"size:255"`
	DueDate     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Order       int       `gorm:"not null;index:,sort:asc,type:btree"`
}

// For controller
type ControllerTask struct {
	ID          uint
	ListID      uint
	Title       string
	Description string
	DueDate     time.Time
	Order       int
}

type SwaggerInputCreateTask struct {
	Title       string `gorm:"not null;size:255;default:Untitled;" json:"Title"`
	Description string `gorm:"size:255" json:"Description"`
	ListID      uint   `gorm:"not null;"  json:"ListID"`
}

type SwaggerInputUpdateTask struct {
	Title       string `gorm:"not null;size:255;default:Untitled;" json:"Title"`
	Description string `gorm:"size:255" json:"Description"`
	ListID      uint   `gorm:"not null;"  json:"ListID"`
	Order       int    `gorm:"not null;index:,sort:asc,type:btree" json:"Order"`
}

type SwaggerTasks struct {
	Tasks []Task `json:"tasks"`
}

type SwaggerTask struct {
	Task Task `json:"task"`
}
