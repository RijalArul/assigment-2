package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string `gorm:"type:varchar(255);not null;unique"`
	OrderedAt    time.Time
	Items        []Item
}
