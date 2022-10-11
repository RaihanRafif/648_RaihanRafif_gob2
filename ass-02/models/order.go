package models

import "time"

type Order struct {
	OrderID      uint   `gorm:"primaryKey"`
	CustomerName string `gorm:"not null;unique;type:varchar(255)"`
	Items        []Item
	OrderedAt    time.Time
	UpdatedAt    time.Time
}
