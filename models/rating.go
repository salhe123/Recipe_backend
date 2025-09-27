package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	UserID   uint
	RecipeID uint
	Score    int `gorm:"check:score >= 1 AND score <= 5"`
}
