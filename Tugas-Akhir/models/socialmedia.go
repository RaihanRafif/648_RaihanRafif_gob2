package models

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" valid:"required"`
	UserID         uint
}
