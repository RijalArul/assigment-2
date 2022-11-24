package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	// ID          uint   `gorm:"primarykey"`
	ItemCode    string `gorm:"not null;" json:"item_code" binding:"required"`
	Description string `gorm:"not null;" json:"desc" binding:"required"`
	Quantity    int    `gorm:"not null;" json:"quantity" binding:"required"`
	OrderID     int    `json:"-" gorm:"foreignKey:OrderID" binding:"required"`
}
