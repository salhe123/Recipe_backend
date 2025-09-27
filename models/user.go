package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Recipes   []Recipe
	Likes     []Like
	Bookmarks []Bookmark
	Comments  []Comment
	Ratings   []Rating
}
