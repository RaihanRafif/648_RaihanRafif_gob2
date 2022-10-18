package models

type Comment struct {
	GormModel
	Message string `json:"message" form:"message"`
	// UserID  User   `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// PhotoID Photo  `json:"photo_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID  uint `json:"user_id" form:"user_id" required:"true"`
	PhotoID uint `json:"photo_id" form:"photo_id" required:"true"`
}
