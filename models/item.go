package models

type Item struct {
	ID          uint   `gorm:"primarykey"`
	ItemCode    string `gorm:"type:varchar(255);not null" json:"item_code"`
	Description string `gorm:"type:varchar(255);not null" json:"desc"`
	Quantity    int    `gorm:"not null" json:"quantity"`
	OrderID     int    `json:"-"`
}
