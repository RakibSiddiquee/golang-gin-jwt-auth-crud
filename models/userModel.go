package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required,min=2,max=50"`
	Email    string `json:"email" gorm:"unique" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Post     Post
	Comment  Comment
}
