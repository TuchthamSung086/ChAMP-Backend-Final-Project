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
	gorm.Model            // https://gorm.io/docs/models.html
	ListID      uint      `gorm:"not null;"`
	Title       string    `gorm:"not null;"`
	Description string    `gorm:"size:255"`
	DueDate     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Order       int       `gorm:"not null;index:,sort:asc,type:btree"`
}

type SwaggerInputCreateTask struct {
	Title       string `gorm:"not null;size:255;" json:"title"`
	Description string `gorm:"size:255" json:"description"`
	ListID      uint   `gorm:"not null;"  json:"list_id"`
}

type SwaggerInputUpdateTask struct {
	Title       string `gorm:"not null;size:255;" json:"title"`
	Description string `gorm:"size:255" json:"description"`
	ListID      uint   `gorm:"not null;"  json:"list_id"`
	Order       int    `gorm:"not null;index:,sort:asc,type:btree" json:"order"`
}

type SwaggerTasks struct {
	Tasks []Task `json:"tasks"`
}

type SwaggerTask struct {
	Task Task `json:"task"`
}
