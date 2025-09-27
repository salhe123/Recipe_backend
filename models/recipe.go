package models

import (
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	PrepTime    int // In minutes
	CategoryID  uint
	UserID      uint
	Category    Category
	Images      []Image
	Steps       []Step
	Ingredients []Ingredient
	Likes       []Like
	Bookmarks   []Bookmark
	Comments    []Comment
	Ratings     []Rating
}
