package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	RecipeID   uint
	URL        string `gorm:"not null"`
	IsFeatured bool   `gorm:"default:false"`
}
