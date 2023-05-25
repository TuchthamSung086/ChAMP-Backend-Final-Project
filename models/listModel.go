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

type SwaggerInputCreateList struct {
	Title string `gorm:"not null;size:255;" json:"title"`
	Order int    `gorm:"not null;index:,sort:asc,type:btree" json:"order"`
}

type SwaggerInputUpdateList struct {
	Title string `gorm:"not null;size:255;" json:"title"`
	Order int    `gorm:"not null;index:,sort:asc,type:btree" json:"order"`
}

type SwaggerLists struct {
	Lists []List `json:"lists"`
}

type SwaggerList struct {
	List List `json:"list"`
}
