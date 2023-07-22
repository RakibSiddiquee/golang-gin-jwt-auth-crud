package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	CategoryID uint   `gorm:"foreignkey:CategoryID" json:"categoryID" binding:"required"`
	Title      string `json:"title" binding:"required,min=2,max=200"`
	Body       string `gorm:"type:text" json:"body"`
	UserID     uint   `gorm:"foreignkey:UserID""`
	Comment    []Comment
}
