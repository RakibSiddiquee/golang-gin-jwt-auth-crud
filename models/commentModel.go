package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID uint   `gorm:"foreignkey:PostID" json:"postID" binding:"required,gt=0"`
	UserID uint   `gorm:"foreignkey:UserID"`
	Body   string `gorm:"type:text"`
}
