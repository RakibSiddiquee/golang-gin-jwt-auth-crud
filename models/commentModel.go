package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID uint   `gorm:"foreignkey:PostID"`
	UserID uint   `gorm:"foreignkey:UserID"`
	Body   string `gorm:"type:text"`
}
