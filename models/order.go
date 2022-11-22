package models

import (
	"time"
)

type Order struct {
	ID           uint   `gorm:"primarykey"`
	CustomerName string `gorm:"type:varchar(255);not null;unique" json:"customer_name"`
	OrderedAt    time.Time
	Items        []Item `json:"items"`
}
