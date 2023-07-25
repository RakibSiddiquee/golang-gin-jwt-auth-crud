package models

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string `gorm:"unique;not null" json:"name"`
	Slug  string `gorm:"unique;not null" json:"slug"`
	Posts []Post
}

func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	category.Slug = slug.Make(category.Name)

	return
}
