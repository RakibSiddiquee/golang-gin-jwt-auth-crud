package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name" binding:"required,min=2"`
	Slug string
	Post []Post
}
