package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	RecipeID uint
	Name     string `gorm:"not null"`
	Quantity string
	Unit     string
}
