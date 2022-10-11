package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Item struct {
	ItemID      uint `gorm:"primaryKey"`
	ItemCode    uint `gorm:"not null;unique;type:varchar(191)"`
	Description string
	Quantity    uint
	OrderID     string
}

func (p *Item) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Item Before Create()")

	if p.Quantity < 1 {
		err = errors.New("Item quantity have bigger than 1")
	}
	return
}
