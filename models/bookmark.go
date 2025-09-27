package models

import "gorm.io/gorm"

type Bookmark struct {
	gorm.Model
	UserID   uint
	RecipeID uint
}
