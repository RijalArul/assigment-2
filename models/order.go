package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	// ID           uint      `gorm:"primarykey"`
	CustomerName string    `gorm:"not null; unique;" json:"customer_name" binding:"required"`
	OrderedAt    time.Time `json:"ordered_at" gorm:"autoCreateTime" binding:"required"`
	Items        []Item    `json:"items" binding:"required"`
}
