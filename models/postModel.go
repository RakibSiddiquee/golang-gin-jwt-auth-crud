package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	CategoryID uint `gorm:"foreignkey:CategoryID"`
	Title      string
	Body       string `gorm:"type:text"`
	UserID     uint   `gorm:"foreignkey:UserID"`
	Comment    Comment
}
