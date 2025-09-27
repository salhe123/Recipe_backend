package models

import "gorm.io/gorm"

type Step struct {
	gorm.Model
	RecipeID    uint
	StepNumber  int    `gorm:"not null"`
	Description string `gorm:"not null"`
}
