package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemCode    string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255);not null"`
	Quantity    int    `gorm:"default:1;not_null;"`
	OrderID     int
}
