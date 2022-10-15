package models

import (
	"tugasakhir/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	UserName string `gorm:"not null" json:"user_name" form:"user_name" valid:"required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
