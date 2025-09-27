package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID   uint
	RecipeID uint
	Content  string `gorm:"not null"`
}
