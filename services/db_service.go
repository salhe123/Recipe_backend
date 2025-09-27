package services

import (
	"fmt"
	"os"
	"recipe_backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return db.AutoMigrate(&models.User{}, &models.Recipe{}, &models.Image{}, &models.Step{},
		&models.Ingredient{}, &models.Category{}, &models.Like{}, &models.Bookmark{},
		&models.Comment{}, &models.Rating{})
}
