package models

type Comment struct {
	GormModel
	Message string `json:"message" form:"message"`
	UserID  uint   `json:"user_id" form:"user_id" required:"true"`
	PhotoID uint   `json:"photo_id" form:"photo_id" required:"true"`
}
