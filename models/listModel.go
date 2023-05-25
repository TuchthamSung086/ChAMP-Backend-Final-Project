package models

import "gorm.io/gorm"

/*
- Description (optional)
- Due date (optional)
- Order
*/

type List struct {
	gorm.Model        // https://gorm.io/docs/models.html
	Title      string `gorm:"not null;size:255;default:Untitled;"`
	Order      int    `gorm:"not null;index:,sort:asc,type:btree"`
	Tasks      []Task `gorm:"foreignKey:ListID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SwaggerInputCreateList struct {
	Title string `gorm:"not null;size:255;default:Untitled;" json:"Title"`
	Order int    `gorm:"not null;index:,sort:asc,type:btree" json:"Order"`
}

type SwaggerInputUpdateList struct {
	Title string `gorm:"not null;size:255;default:Untitled;" json:"Title"`
	Order int    `gorm:"not null;index:,sort:asc,type:btree" json:"Order"`
}

type SwaggerLists struct {
	Lists []List `json:"lists"`
}

type SwaggerList struct {
	List List `json:"list"`
}
