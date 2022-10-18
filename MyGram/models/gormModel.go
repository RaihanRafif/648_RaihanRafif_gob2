package models

type GormModel struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
}
