package models

import "gorm.io/gorm"

/*
- Description (optional)
- Due date (optional)
- Order
*/

type List struct {
	gorm.Model        // https://gorm.io/docs/models.html
	Title      string `gorm:"not null;size:255;"`
	Order      int    `gorm:"not null;index:,sort:asc,type:btree"`
	Tasks      []Task `gorm:"foreignKey:ListID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SwaggerList struct {
	List []List `json:"list"`
}
