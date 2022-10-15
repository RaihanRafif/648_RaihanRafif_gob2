package models

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title"`
	Caption  string `json:"caption" form:"caption" valid:"required"`
	PhotoUrl string `json:"photo_url" form:"photo_url" valid:"required"`
	UserID   uint   `json:"user_id" form:"user_id" valid:"required"`
}
