package models

type Comment struct {
	GormModel
	Message string `json:"message" form:"message"`
	UserID  uint
	PhotoID uint
}
