package models

import "gorm.io/gorm"

/*
- Description (optional)
- Due date (optional)
- Order
*/

type List struct {
	gorm.Model        // https://gorm.io/docs/models.html
	Title      string `gorm:"not null;"`
	Order      int    `gorm:"not null;"`
}
